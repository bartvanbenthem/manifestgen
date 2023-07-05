use std::collections::BTreeMap;
use std::fs;
use std::env;
use std::error::Error;
use std::io::{self, Read, Write};
use std::process;

use serde::{Deserialize, Serialize};
use serde_json::{Value};
use serde_yaml;
use handlebars::Handlebars;
use clap::{App, Arg};

#[derive(Debug, Deserialize, Serialize)]
struct KeyValuePairs(BTreeMap<String, Value>);

#[derive(Debug)]
struct Config {
    template_file: String,
    output_file: String,
    variables_file: String,
}

fn main() ->  Result<(), Box<dyn Error>> {
    // Get command-line arguments
    let config = get_args().unwrap();

    let input_type_result = parse_input_type(&config);
    let input_type = match input_type_result {
        Ok(input_type) => input_type,
        Err(error) => {
            println!("{}",error);
            eprintln!("No valid JSON or YAML input is given, restart the Manifestgen with valid input!");
            process::exit(1)
        },
    };

    let rendered_template_result = render_config(&input_type, &config);
    let rendered_template = match rendered_template_result {
        Ok(rendered_template) => rendered_template,
        Err(error) => {
            eprintln!("Error while rendering {}", error);
            process::exit(1)
        },
    };

    // write to disk or stdout based on the provided output param
    if let Err(err) = manifest_writer(&config.output_file, &rendered_template) {
        eprintln!("Error: {}", err);
        process::exit(1)
    }

    Ok(())
}

// --------------------------------------------------
fn get_args() -> Result<Config, Box<dyn Error>> {
    // Define and parse command-line arguments using clap
    let matches = App::new("Template Renderer")
        .arg(
            Arg::with_name("template_file")
                .short("t")
                .long("template")
                .required(true)
                .takes_value(true)
                .help("Path to the template file"),
        )
        .arg(
            Arg::with_name("output_file")
                .short("o")
                .long("output")
                .required(false)
                .takes_value(true)
                .help("Path to the output file"),
        )
        .arg(
            Arg::with_name("variables_file")
                .short("v")
                .long("variables")
                .required(false)
                .takes_value(true)
                .help("Path to the JSON file"),
        )
        .get_matches();

    Ok(Config {
        template_file: matches.value_of("template_file").unwrap().to_string(),
        output_file: matches.value_of("output_file").unwrap_or("").to_string(),
        variables_file: matches.value_of("variables_file").unwrap_or("").to_string(),
    })
}

// --------------------------------------------------
fn parse_input_type(config: &Config) -> Result<KeyValuePairs, Box<dyn Error>> {
         // initialize KeyValuePair type
         let key_value_pairs: KeyValuePairs;

         if config.variables_file.is_empty() {
             // Process the JSON value trough stdin
             let mut input = String::new();
             io::stdin().read_to_string(&mut input)?;
             if is_valid_json(&input) {
                key_value_pairs = serde_json::from_str(&input)?;
             } else if is_valid_yaml(&input){
                key_value_pairs = serde_yaml::from_str(&input)?;
             } else {
                panic!("No valid JSON or YAML input type!")
             }
         } else {
             // Read JSON file
             let var_file = fs::read_to_string(&config.variables_file)?;
             if is_valid_json(&var_file) {
                key_value_pairs = serde_json::from_str(&var_file)?;
             } else if is_valid_yaml(&var_file){
                key_value_pairs = serde_yaml::from_str(&var_file)?;
             } else {
                panic!("No valid JSON or YAML input type!")
             }
         }

         Ok(key_value_pairs)

}

fn is_valid_json(input: &str) -> bool {
    match serde_json::from_str::<Value>(input) {
        Ok(_) => true,
        Err(_) => false,
    }
}

fn is_valid_yaml(input: &str) -> bool {
    match serde_yaml::from_str::<Value>(input) {
        Ok(_) => true,
        Err(_) => false,
    }
}

// --------------------------------------------------
fn render_config(key_value_pairs: &KeyValuePairs, config: &Config) -> Result<String, Box<dyn Error>> {
        // Load template file
        let template_content = fs::read_to_string(&config.template_file)?;
    
        // Initialize the templating engine
        let mut handlebars = Handlebars::new();
        handlebars.register_template_string("template", &template_content)?;
    
        // Render the template
        let rendered_template = handlebars.render("template", &key_value_pairs.0)?;

        Ok(rendered_template)
}

// --------------------------------------------------
fn manifest_writer(output: &String, template: &String) -> Result<(), Box<dyn Error>> {
    if output.is_empty() {
        // Get the stdout handle
        let stdout = io::stdout();
        let mut handle = stdout.lock();
        // Write the JSON string to stdout
        handle.write_all(template.as_bytes())?;
        // Make sure the output is flushed
        handle.flush()?;
        println!("");
    } else {
        // Print the current working directory
        print_current_dir();
        // Write the rendered template to the output file
        fs::write(output, template)?;
        println!("Template successfully rendered and written to {}!", output);
    }    

    Ok(())
}

// --------------------------------------------------
fn print_current_dir() {
    if let Ok(current_dir) = env::current_dir() {
        println!("Current Directory: {}", current_dir.display());
    } else {
        println!("Failed to retrieve the current directory");
    }
}


// --------------------------------------------------
// UNIT TESTS
// --------------------------------------------------

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::{self, Write};
    use std::env;
    use std::sync::{Mutex, Arc};

    // Custom writer that collects the output into a buffer
    struct BufferWriter(Arc<Mutex<Vec<u8>>>);

    impl Write for BufferWriter {
        fn write(&mut self, buf: &[u8]) -> io::Result<usize> {
            self.0.lock().unwrap().write(buf)
        }

        fn flush(&mut self) -> io::Result<()> {
            self.0.lock().unwrap().flush()
        }
    }

    #[test]
    fn test_manifest_writer_with_non_empty_output() {
        // Create a temporary directory
        let temp_dir = env::temp_dir();

        // Generate a path for the output file within the temporary directory
        let output_file = temp_dir.join("output.json");

        // Convert the output file path to a string
        let output = output_file.to_string_lossy().into_owned();

        // Set up test inputs
        let template = "{\"key\": \"value\"}".to_string();

        // Call the function being tested
        manifest_writer(&output, &template).unwrap();

        // Read the content of the output file
        let file_content = fs::read_to_string(output_file).unwrap();

        // Define the expected file content
        let expected_content = "{\"key\": \"value\"}";

        // Assert that the file content matches the expected content
        assert_eq!(file_content, expected_content);
        
        // Cleanup: Delete the temporary output file
        fs::remove_file(&output).unwrap();
    }
}

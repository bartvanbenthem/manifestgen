use std::collections::BTreeMap;
use std::fs;
use serde::{Deserialize, Serialize};
use serde_json;
use std::env;
use std::error::Error;
use std::io::{self, Read, Write};

use handlebars::Handlebars;
use clap::{App, Arg};

#[derive(Debug, Deserialize, Serialize)]
struct KeyValuePairs(BTreeMap<String, String>);

#[derive(Debug)]
struct Config {
    template_file: String,
    output_file: String,
    json_file: String,
    yaml_file: String,
}

fn main() ->  Result<(), Box<dyn Error>> {
    // Get command-line arguments
    let config = get_args().unwrap();
    
    // initialize KeyValuePair type
    let key_value_pairs: KeyValuePairs;

    if config.json_file.is_empty() {
        // Process the JSON value trough stdin
        let mut input = String::new();
        io::stdin().read_to_string(&mut input)?;
        key_value_pairs = serde_json::from_str(&input)?;
    } else {
        // Read JSON file
        let json_content = fs::read_to_string(config.json_file)?;
        key_value_pairs = serde_json::from_str(&json_content)?;
    }
 
    // Load template file
    let template_content = fs::read_to_string(config.template_file)?;

    // Initialize the templating engine
    let mut handlebars = Handlebars::new();
    handlebars.register_template_string("template", &template_content)?;

    // Render the template
    let rendered_template = handlebars.render("template", &key_value_pairs.0)?;

    // write to disk or stdout based on the provided output param
    if let Err(err) = manifest_writer(&config.output_file, &rendered_template) {
        eprintln!("Error: {}", err);
        // Handle the error here
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
            Arg::with_name("variables_json")
                .short("j")
                .long("variables-json")
                .required(false)
                .takes_value(true)
                .help("Path to the JSON file"),
        )
        .arg(
            Arg::with_name("variables_yaml")
                .short("y")
                .long("variables-yaml")
                .required(false)
                .takes_value(true)
                .help("Path to the YAML file"),
        )
        .get_matches();

    Ok(Config {
        template_file: matches.value_of("template_file").unwrap().to_string(),
        output_file: matches.value_of("output_file").unwrap_or("").to_string(),
        json_file: matches.value_of("variables_json").unwrap_or("").to_string(),
        yaml_file: matches.value_of("variables_yaml").unwrap_or("").to_string(),
    })
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

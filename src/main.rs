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

type MyResult<T> = Result<T, Box<dyn Error>>;

#[derive(Debug)]
struct Config {
    template_file: String,
    output_file: String,
    json_file: String,
}

fn main() ->  MyResult<()> {
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

fn get_args() -> MyResult<Config> {
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
        json_file: matches.value_of("variables_json").unwrap_or("").to_string(),
    })
}

fn manifest_writer(output: &String, template: &String) -> Result<(), Box<dyn std::error::Error>> {
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

fn print_current_dir() {
    if let Ok(current_dir) = env::current_dir() {
        println!("Current Directory: {}", current_dir.display());
    } else {
        println!("Failed to retrieve the current directory");
    }
}
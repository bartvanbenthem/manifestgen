use std::process;

fn main() {
    // Get command-line arguments
    let config = manifestgen::get_args().unwrap();

    let input_type_result = manifestgen::parse_input_type(&config);
    let input_type = match input_type_result {
        Ok(input_type) => input_type,
        Err(err) => {
            println!("{}", err);
            eprintln!(
                "No valid JSON or YAML input is given, restart the Manifestgen with valid input!"
            );
            process::exit(1)
        }
    };

    // render manifests
    let rendered_template_result = manifestgen::render_config(&input_type, &config);
    let rendered_template = match rendered_template_result {
        Ok(rendered_template) => rendered_template,
        Err(err) => {
            eprintln!("Error while rendering {}", err);
            process::exit(1)
        }
    };

    // write to disk or stdout based on the provided output param
    if let Err(err) = manifestgen::manifest_writer(&config.output_file, &rendered_template) {
        eprintln!("Error: {}", err);
        process::exit(1)
    }
}
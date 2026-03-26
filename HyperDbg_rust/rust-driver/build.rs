use wdk_build::Config;

fn main() -> Result<(), wdk_build::Error> {
    let config = Config::from_env_auto()?;
    wdk_build::build_cdylib(config)?;
    Ok(())
}

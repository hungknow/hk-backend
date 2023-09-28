use economy_core::Ohlc;

use super::OhlcCsvParser;

// "<TICKER>,<DTYYYYMMDD>,<TIME>,<OPEN>,<HIGH>,<LOW>,<CLOSE>,<VOL>"
pub struct OhlcCsvParserFile {}

impl OhlcCsvParser for OhlcCsvParserFile {
    fn parse() -> Vec<Ohlc> {
        // Detect the header line
        // Based on the header, create the column
        Vec::new()
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn detect_header() {

    }
}
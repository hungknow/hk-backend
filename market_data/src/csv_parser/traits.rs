
use economy_core::Ohlc;

pub trait OhlcCsvParser {
    fn parse() -> Vec<Ohlc>;
}

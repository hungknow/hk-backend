use crate::Ohlc;

pub trait OhlcCsvParser {
    fn parse() -> Vec<Ohlc>;
}

pub trait OhlcReader {
    fn get(&self, symbol: &str, query: &str) -> Vec<Ohlc>;
}

pub struct OhlcReplayState {
    pub read_index: u32,
    pub available_count: u32,
}

/**
 * Activity:
 * - Returns one next candle bar
 * - Subtract the latest candle bar
 * - Get all candle bars get so far
 * - Get the read index
 * - Get the available indexes (optional)
 */
pub trait OhlcReplay {
    fn get_state(&self) -> OhlcReplayState;
    fn set_read_index(&self, new_read_index: u32);
    fn get_ohlc_up_to_read_index(&self, read_index: u32) -> Vec<Ohlc>;
    
    // fn increase_read_index(&self, interval: u32)
}

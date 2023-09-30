use std::io::ErrorKind;

use chrono::{DateTime, Utc};
use economy_core::{OhlcReplay, OhlcReplayState, EconomyError, Ohlc};

pub struct FakeOhlcReplay {
    state: OhlcReplayState, // Maximum state
}

impl FakeOhlcReplay {
    fn new() -> FakeOhlcReplay {
        FakeOhlcReplay {
            state: OhlcReplayState {
                read_index: 0,
                available_count: 10,
            },
        }
    }
}

impl OhlcReplay for FakeOhlcReplay {
    fn get_state(&self) -> economy_core::OhlcReplayState {
        self.state
    }

    fn set_read_index(&self, new_read_index: u32) -> Result<u32, EconomyError> {
        if new_read_index < 0 || new_read_index > self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }
        let old_read_index = self.state.read_index;
        self.state.read_index = new_read_index;
        Ok(old_read_index) 
    }

    fn get_ohlc_up_to_read_index(&self, read_index: u32) -> Result<Vec<economy_core::Ohlc>, EconomyError> {
        if read_index < 0 || read_index > self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }
        let mut vec = Vec::new();
        vec.push(Ohlc{
            close_time: chrono::offset::Utc::now(),
            high: 0.0,
            open: 0.0,
            close: 0.0,
            low: 0.0,
            volume: Some(0.0),
            trade_count: Some(0.0),
        });
        Ok(vec)
    }
}

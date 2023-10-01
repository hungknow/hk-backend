use std::ops::Sub;

use economy_core::{EconomyError, Ohlc, OhlcReplay, OhlcReplayState};

pub struct FakeOhlcReplay {
    state: OhlcReplayState, // Maximum state
    vec: Vec<Ohlc>,
}

impl FakeOhlcReplay {
    pub fn new() -> FakeOhlcReplay {
        let target_available_count: u32 = 10;
        let mut vec = Vec::new();
        for n in 1..target_available_count + 1 {
            vec.push(Ohlc {
                close_time: chrono::Utc::now().sub(chrono::Duration::seconds(i64::from(
                    target_available_count - n,
                ))),
                high: 0.0,
                open: 0.0,
                close: 0.0,
                low: 0.0,
                volume: Some(0.0),
                trade_count: Some(0.0),
            });
        }

        FakeOhlcReplay {
            state: OhlcReplayState {
                read_index: 0,
                available_count: target_available_count,
            },
            vec: vec,
        }
    }
}

impl OhlcReplay for FakeOhlcReplay {
    fn get_state(&self) -> economy_core::OhlcReplayState {
        self.state
    }

    fn set_read_index(&mut self, new_read_index: u32) -> Result<u32, EconomyError> {
        if new_read_index > self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }
        let old_read_index = self.state.read_index;
        self.state.read_index = new_read_index;
        Ok(old_read_index)
    }

    fn get_ohlc_up_to_read_index(
        &mut self,
        read_index: u32,
    ) -> Result<Vec<economy_core::Ohlc>, EconomyError> {
        if read_index > self.state.available_count {
            return Err(EconomyError::OutOfRange);
        }
        Ok(self.vec.clone())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ohlc_replay() {
        let mut fake_ohlc_replay = FakeOhlcReplay::new();
        let state = fake_ohlc_replay.get_state();
        assert_eq!(state.available_count, 10);
        assert_eq!(state.read_index, 0);

        assert_eq!(fake_ohlc_replay.set_read_index(1), Ok(0));
    }
}

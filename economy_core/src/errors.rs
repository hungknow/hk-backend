use core::fmt;

#[derive(Debug, PartialEq)]
pub enum EconomyError {
    OutOfRange,
    InvalidParameter,
}

impl fmt::Display for EconomyError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}", self)
    }
}

pub type Result<T> = std::result::Result<T, EconomyError>;

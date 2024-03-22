use core::fmt;
use csv::Error as CsvError;

#[derive(Debug)]
pub enum EconomyError {
    OutOfRange,
    InvalidParameter,
    Csv(CsvError),
}

impl fmt::Display for EconomyError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}", self)
    }
}

// impl Error for EconomyError {
//     fn source(&self) -> Option<&(dyn Error + 'static)> {
//        match *self {

//        } 
//     }
// }

impl From<csv::Error> for EconomyError {
    fn from(err: csv::Error) -> Self {
        EconomyError::Csv(err)
    }
}

pub type Result<T> = std::result::Result<T, EconomyError>;

# utils.py

from dataclasses import dataclass
from datetime import datetime
import re
from typing import List, Optional

@dataclass(frozen=True)
class DateParser:
    date_string: str

    @classmethod
    def parse(cls, date_string: str) -> 'DateParser':
        """
        Parse a date string in the format 'YYYY-MM-DD' into a DateParser object.
        
        Args:
        date_string (str): The date string to parse.
        
        Returns:
        DateParser: A DateParser object.
        """
        try:
            date = datetime.strptime(date_string, '%Y-%m-%d')
            return cls(date_string=date_string)
        except ValueError:
            raise ValueError(f"Invalid date string: {date_string}")

    @property
    def year(self) -> int:
        return self.date_string.split('-')[0]

    @property
    def month(self) -> int:
        return int(self.date_string.split('-')[1])

    @property
    def day(self) -> int:
        return int(self.date_string.split('-')[2])

def clean_text(text: str) -> str:
    """
    Clean a string by removing leading and trailing whitespace and replacing multiple spaces with a single space.
    
    Args:
    text (str): The string to clean.
    
    Returns:
    str: The cleaned string.
    """
    return re.sub(r'\s+', ' ', text).strip()

def extract_ids(text: str) -> List[int]:
    """
    Extract all integer IDs from a string.
    
    Args:
    text (str): The string to extract IDs from.
    
    Returns:
    List[int]: A list of integer IDs.
    """
    return [int(i) for i in re.findall(r'\d+', text)]
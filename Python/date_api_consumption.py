import requests


def parse_date(month:str,day:str) -> str:
    """
    Convert the Month name and day number to a string of format m/dd for
    consumption of the API.  Even if a number such as 85 is put in, the API
    will convert it to some valid date.  So one could put in 2/31 and the API would convert this
    to March 3rd for a non-leapyear year.

    :param month: A name of a month e.x January or even january will work.
    :type month: str
    :param day: A day number for example 31.
    :type day: str
    :raises Exception: If the month doesn't exist then an error is thrown.
    :return: Return the format "m/dd"
    :rtype: str
    """
    month = month.title()
    month_map = {
        'January': 1, 'February': 2, 'March': 3,
        'April': 4, 'May': 5, 'June': 6,
        'July': 7, 'August': 8, 'September': 9,
        'October': 10, 'November': 11, 'December': 12
    }
    month_number = month_map.get(month,-1)
    if month_number == -1:
        raise Exception(f"Month {month} is not a valid month.  Try one of the following {list(month_map.keys())}")
    
    return f"{month_number}/{day}"



def get_random_date_fact(date_str:str) -> str:
    """
    Consumes the numbersapi.com random facts api

    :param date_str: The parsed date_str from parse_date function.
    :type date_str: str
    :return: Returns the random fact about that date.
    :rtype: str
    """
    url = f"http://numbersapi.com/{date_str}/date"
    response = requests.get(url)
    return response.text

def main():

    month = input("Input a month to get a random fact for and press [Enter] (e.g. January): ")
    day = input("Input the day of the month to get a random fact for and press [Enter] (e.g. 27): ")

    date_str = parse_date(month,day)
    random_fact = get_random_date_fact(date_str)

    print()
    print()
    print(f"==================Random Fact about {date_str}==================")
    print()
    print()
    print(random_fact)


if __name__ == "__main__":
    main()


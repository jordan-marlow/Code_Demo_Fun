import requests



def get_weather_information(city:str) -> dict:
    """
    This grabs the information of a city from the geoweather API

    :param city: The desired city to get information for.
    :type city: str
    :return: The dictionary of information that is retrieved.
    :rtype: dict
    """
    url = f"https://goweather.herokuapp.com/weather/{city}"
    session = requests.Session()
    r = session.get("https://reacttempo.netlify.app/")
    print(r)
    response = session.get(url)
    print(response)
    try:
        data = response.json()
    except requests.exceptions.JSONDecodeError:
        data = {}
    return data

def parse_temperature(temp:str) -> float:
    """
    The temerature from the API comes in as "+XY °C".
    We only want "+/-XY" to be returned as a float so we can
    convert or use it as a number.
    :param temp: temperature parameter from the goweather API
    :type temp: str
    :return: The value of the temperature in Celsius as a float
    :rtype: float
    """
    if not temp:
        return -1
    temp = temp.split(" ")[0]
    sign = temp[0]
    if sign == "+":
        return float(temp[1:])
    return float(temp[1:])*-1

def parse_speed(speed:str) -> float:
    """
    The wind speed form the API comes in as "XY km/h".
    We only want the "XY" as a float.

    :param speed: wind speed from the geoweather API
    :type speed: str
    :return: The value of the wind speed in km/h as a float.
    :rtype: float
    """
    if not speed:
        return -1
    speed = speed.split(' ')[0]
    return float(speed)


def convert_C_to_F(temp:float) -> float:
    """
    Convert Celsius to Fahrenheit

    :param temp: The temperature in Celsius to be converted
    :type temp: float
    :return: The temperature in Fahrenheit
    :rtype: float
    """
    freedom_temp = temp*(9/5) + 32
    return freedom_temp

def convert_kmh_to_mph(speed:float) -> float:
    """
    Convert km/h to mph

    :param speed: The speed in km/h to be converted
    :type speed: float
    :return: The speed in mph
    :rtype: float
    """
    mph = round(speed/1.609,2)
    return mph

def main():
    city = input("Enter a city to get the weather for and press [Enter]:  ")
    data = get_weather_information(city)
    print(data)
    temp_celsius = parse_temperature(data.get("temperature",""))
    wind_kmh = parse_speed(data.get("wind",""))
    description = data.get("description","")

    temp_fahr = convert_C_to_F(temp_celsius)
    wind_mph = convert_kmh_to_mph(wind_kmh)

    print(f"City:\t{city}")
    print()
    print()
    print(f"Description:   {description}")
    print(f"Temperature:   {temp_celsius}°C     {temp_fahr}°F")
    print(f"Wind Speed:    {wind_kmh} km/h   {wind_mph} mph")

if __name__ == "__main__":
    main()
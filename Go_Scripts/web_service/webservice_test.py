import requests
import json


# Test Dictionary for microservices
base_url = "http://localhost:8080/api"
urls = {"get_patients":{"num_patients":8},
        "cc_validate":{"card_number":"4532015112830366"}}


# Send POST request
for url in urls:
    gather_url = base_url + "/" + url
    response = requests.get(gather_url, data=json.dumps(urls[url]), headers={"Content-Type": "application/json"})

    # Print the response from the server
    print("Status Code:", response.status_code)
    print("Response JSON:", response.json())

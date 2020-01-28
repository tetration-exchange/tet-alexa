from tetpyclient import RestClient
import json

secret = "17e49e0f84a1acce67998c33a61a298fb8dfd98c"
key = "b21a1efff5b64e79a970222b15e7bc4c"

tetration = RestClient("https://demo.tetrationpreview.com",
                       api_key=key,
                       api_secret=secret)

workloads = "5dc374a4755f0278711e3a0f"
quarantine_tag = "5e2ebe52755f025d31d980b5"

req_payload = {
    "version": "v0",
    "rank": "ABSOLUTE",
    "policy_action": "DENY",
    "priority": 99,
    "consumer_filter_id": quarantine_tag,
    "provider_filter_id": workloads
}

resp = tetration.post('/applications/5dc60ac8755f0260fd1e3a0d/policies',
                      json_body=json.dumps(req_payload))

print "Created the policy with id:", resp.json()["id"]

policy_id = resp.json()["id"]

req_payload = {"version": "v0", "proto": None}

resp = tetration.post('/policies/{}/l4_params'.format(policy_id),
                      json_body=json.dumps(req_payload))

print "Created the service ports with id:", resp.json()["id"]
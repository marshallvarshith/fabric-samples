from hfc.fabric import Client as FabricClient
import json

class AccessLogChaincode:

    def __init__(self):
        self.client = FabricClient()

    def initLedger(self):
        return 'Ledger initialized'

    def createLog(self, user_id, access_method, timestamp):
        # Check if key already exists
        if self.client.get_state(user_id):
            return f'Error: Key {user_id} already exists'

        # Create new log
        log_entry = {
            'user_id': user_id,
            'access_method': access_method,
            'timestamp': timestamp
        }
        self.client.put_state(user_id, json.dumps(log_entry))
        return f'Log created for {user_id}'

    def queryLog(self, user_id):
        log_entry = self.client.get_state(user_id)
        if not log_entry:
            return f'Error: No log found for {user_id}'

        return json.loads(log_entry)

def main():
    chaincode = AccessLogChaincode()
    # Example usage
    print(chaincode.initLedger())
    print(chaincode.createLog('user1', 'manual', '2024-07-22T12:00:00Z'))
    print(chaincode.queryLog('user1'))

if __name__ == '__main__':
    main()

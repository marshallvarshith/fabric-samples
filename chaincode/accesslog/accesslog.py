from hfc.fabric import Client as FabricClient
from hfc.fabric.chaincode import Chaincode
import json

class AccessLogChaincode(Chaincode):
    def __init__(self):
        super().__init__('accesslog')

    def create_access_log(self, stub, args):
        user_id, access_method, timestamp = args
        access_log = {
            'userID': user_id,
            'accessMethod': access_method,
            'timestamp': timestamp
        }
        key = f"{user_id}_{timestamp}"
        stub.put_state(key, json.dumps(access_log).encode('utf-8'))
        return key

    def query_access_log(self, stub, args):
        key = args[0]
        access_log = stub.get_state(key)
        if not access_log:
            return f"Access log {key} not found"
        return access_log.decode('utf-8')

cc = AccessLogChaincode()

if __name__ == '__main__':
    cc.start()

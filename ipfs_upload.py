from ipfshttpclient.client import connect
import io
import ipfshttpclient
#from storage_methods.abstract_storage import Storage


class IPFSStorage():
    def __init__(self, connection_string: str):
        self.ip = ip
        self.port = port
        self.client = ipfshttpclient.connect(connection_string)

    def store(self, filename,content):
        multihash = self.client.add(io.BytesIO(content))
        return multihash["Hash"]
    def retrieve(self,multihash):
        return self.client.get(multihash)
    def remove(self,filename):
        return



if __name__ == "__main__":
    ipfs_storage_1 = IPFSStorage("/ip4/128.199.19.171/tcp/5001")
    ipfs_storage_2 = IPFSStorage("/ip4/128.199.19.171/tcp/5001")
    hash_1_a = ipfs_storage_1.store("a")
    hash_1_b = ipfs_storage_1.store("b")
    hash_1_c = ipfs_storage_1.store("c")
    hash_1_d = ipfs_storage_1.store("d")
    hash_2_a = ipfs_storage_2.store("a")
    hash_2_b = ipfs_storage_2.store("b")
    hash_2_c = ipfs_storage_2.store("c")
    hash_2_d = ipfs_storage_2.store("d")
    # write hashes to file
    with open("hashes.txt", "w") as f:
        f.write(hash_1_a + "\n")
        f.write(hash_1_b + "\n")
        f.write(hash_1_c + "\n")
        f.write(hash_1_d + "\n")
        f.write(hash_2_a + "\n")
        f.write(hash_2_b + "\n")
        f.write(hash_2_c + "\n")
        f.write(hash_2_d + "\n")





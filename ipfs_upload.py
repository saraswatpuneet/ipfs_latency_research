import ipfsApi

class IPFSStorage():
    def __init__(self, ip:str, port=5001):
        self.client = ipfsApi.Client(ip, port)
    def store(self, filename):
        with open(filename, 'rb') as f:
            multihash = self.client.add(f)
            return multihash["Hash"]

if __name__ == "__main__":
    ipfs_storage_1 = IPFSStorage("<ip>")
    ipfs_storage_2 = IPFSStorage("<ip>")
    hash_test = ipfs_storage_1.store("test.txt")
    #hash_1_a = ipfs_storage_1.store("a")
    #hash_2_a = ipfs_storage_2.store("a")
    #hash_1_b = ipfs_storage_1.store("b")
    #hash_2_b = ipfs_storage_2.store("b")
    #hash_1_c = ipfs_storage_1.store("c")
    #hash_2_c = ipfs_storage_2.store("c")
    hash_1_d = ipfs_storage_1.store("d")
    hash_2_d = ipfs_storage_2.store("d")
    # write hashes to file
    with open("hashes_d.txt", "w") as f:
        #f.write(hash_1_a + "\n")
        #f.write(hash_1_b + "\n")
        #f.write(hash_1_c + "\n")
        f.write(hash_1_d + "\n")
        #f.write(hash_2_a + "\n")
        #f.write(hash_2_b + "\n")
        #f.write(hash_2_c + "\n")
        f.write(hash_2_d + "\n")

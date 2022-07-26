import ipfsApi

class IPFSStorage():
    def __init__(self, ip:str, port=5001):
        self.client = ipfsApi.Client(ip, port)
    def store(self, filename):
        with open(filename, 'rb') as f:
            multihash = self.client.add(f)
            return multihash["Hash"]
    def retrieve(self,multihash):
        return self.client.get(multihash)


if __name__ == "__main__":
    ipfs_storage_1 = IPFSStorage("ip1")
    ipfs_storage_2 = IPFSStorage("ip2")
    hash_test = ipfs_storage_1.store("test.txt")
    print(hash_test)
    print(ipfs_storage_1.retrieve(hash_test))
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
    # read hashes from file
    with open("hashes.txt", "r") as f:
        hashes = f.readlines()
    # retrieve hashes from storage
    for hash in hashes:
        print(ipfs_storage_1.retrieve(hash))

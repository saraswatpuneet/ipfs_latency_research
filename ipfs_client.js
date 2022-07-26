import {create} from 'ipfs-http-client';
import {readFileSync} from 'fs';
const ipfs1 = create({ host: "ip", port: 5001, protocol: "http" });
const ipfs2 = create({ host: "ip", port: 5001, protocol: "http" });
const addFile1 = async (fileName, filePath) => {
    // timer 
    const start = Date.now();
    const file = readFileSync(filePath);
    const filesAdded = await ipfs1.add({ path: fileName, content: file }, {
    progress: (len) => console.log("Uploading file..." + len)
  });
    console.log(filesAdded);
    const fileHash = filesAdded.cid.string;
    // get file from ipfs
    const fileFromIpfs = await ipfs1.cat(fileHash);
    let content = []
    for await(const chunk of fileFromIpfs) {
        content.push(chunk);
    }
    const raw = Buffer.from(content).toString('utf8')
    console.log(JSON.parse(raw))
    // timer
    const end = Date.now();
    const time = end - start;
    console.log("Time taken to upload file: ", time);
    return fileHash;
};

const addFile2 = async (fileName, filePath) => {
    // timer 
    const start = Date.now();
    const file = readFileSync(filePath);
    const filesAdded = await ipfs2.add({ path: fileName, content: file }, {
    progress: (len) => console.log("Uploading file..." + len)
  });
    console.log(filesAdded);
    const fileHash = filesAdded.cid.string;
    // get file from ipfs
    const fileFromIpfs = await ipfs2.cat(fileHash);
    let content = []
    for await(const chunk of fileFromIpfs) {
        content.push(chunk);
    }
    const raw = Buffer.from(content).toString('utf8')
    console.log(JSON.parse(raw))
    const end = Date.now();
    const time = end - start;
    console.log("Time taken to upload file: ", time);
    return fileHash;
};

addFile1("test.txt", "./test.txt");
addFile2("test.txt", "./test.txt");
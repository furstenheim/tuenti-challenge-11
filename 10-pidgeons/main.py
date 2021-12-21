import scapy.all as scapy
import base64

capture = scapy.rdpcap('icmps.pcap')
ping_data = b""

result = []
for packet in capture:
    if packet[scapy.ICMP].type == 8:
        result.append({
            'data': packet.load,
            ## also seq
            'id': packet[scapy.ICMP].fields['seq']
        })
    else:
        print("different type")

result.sort(key=lambda el: el['id'])
print(result)


for package in result:
    ping_data += package['data']

with open("result.txt", "wb") as writeFile:
    writeFile.write(ping_data)
print(ping_data)
# print(base64.b64decode(ping_data))

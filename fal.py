import binascii
import sys
import falcon

N = 128
M = "hello"
if (len(sys.argv) > 1):
    M = str(sys.argv[1])
if (len(sys.argv) > 2):
    N = int(sys.argv[2])

print("Message: ", M)
M = M.encode()

sk = falcon.SecretKey(N)
pk = falcon.PublicKey(sk)
print("Private key:", sk)
print("Public key: ", pk)

sig = sk.sign(M)
print("\nSignature: ", binascii.hexlify(sig))
rtn = pk.verify(M, sig)
print("\nSignature verification: ", rtn)

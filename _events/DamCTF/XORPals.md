---
ctf: DamCTF 2021
points: 188
solved: true
title: xorpals
category: crypto
difficulty: none
flag: dam{antman_EXPANDS_inside_tHaNoS_never_sinGLE_cHaR_xOr_yeet}
---



> One of the 60-character strings in the provided file has been encrypted by single-character XOR. 
> The challenge is to find it, as that is the flag.
>
> **Hint**: Always operate on raw bytes, never on encoded strings. Flag must be submitted as UTF8 string.

Attachments: **flags.txt**

```bash
➜  cryptoxor cat flags.txt | head -5
045c3f704f355f6e70536d1e4246573c34096b022f1a077a1d2b676052275f493618787c5a250545254e12750c2261511e5c0d0045376722002a6602
6c3c73194c4e3e3a0563684b600b5c7f1333044e622534244065241e1e0f5f515245546d2030455518065220006b0e3c4b621064732340721f332225
31182e4e4f635b4d506c54282b764a6e70763f24755b1b694a2e0e2c070d0201397277511e72762b2f3a2037720b442e143f5b706f7602787c22643f
52681b5d39262d102e420b42545a085f28581e401f6f657d2e0b5f35357b1569787572466b4f5b106a7975371f537a137c2b671e7972327d4d2b7f4a
177e19705e55251f704e7632796e772728374a63382d5d314b390849747728496d09101458682a2e587400124845677e5d24174c1c0a64396e24091e
```

### flags.txt

flags.txt contained 99 seemingly random hex strings.

# Solution

The challenge said that one of these 99 hex strings was encrypted by single-character XOR. We can solve this in many ways. The easiest is to go through all of the 99 hex strings, with characters (I chose bytes ranging from 0 to 122, since z is the last letter in the ASCII table). We also know the flag begins with `dam{`, so we check for decrypted strings that start with `dam{`.

```python
from pwn import xor
import binascii

with open("flags.txt", "r") as f:
    lines = []
    for i in range(98):
        lines.append(binascii.unhexlify(f.readline()[:-1]))
    for i in range (122):
        for l in lines:
            # pass
            x = xor(l, chr(i))
            if x.startswith(b"dam{"):
                print(x.decode("utf8"))
```

### Output:

dam{antman_EXPANDS_inside_tHaNoS_never_sinGLE_cHaR_xOr_yeet}

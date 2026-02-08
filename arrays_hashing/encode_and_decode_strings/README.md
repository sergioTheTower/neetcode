# Encode and Decode Strings

## 📝 Problem Description

Design an algorithm to encode a list of strings to a single string. The encoded string is then sent over a network and must be decoded back to the original list of strings.

**Machine 1 (Sender)** implements:
`string encode(vector<string> strs)`

**Machine 2 (Receiver)** implements:
`vector<string> decode(string s)`

### Constraints

- `0 <= strs.length < 100`
- `0 <= strs[i].length < 200`
- `strs[i]` contains any of the 256 valid ASCII characters.

---

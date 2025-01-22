# SimpleFinancePackage


An Package to interact with an mock stock exchange

## Structure
| Field | Length |
| ----- | ------ |
| Key   | 8 Bytes |
| Operation | Byte |
| Asset | 2 Bytes |
| Price | Float64 |
| Volume | Float64 |
| Expire (Unix) | uint64 |
| CRC | 2 Bytes |

## TODOS
- [ ] Test it  
- [ ] Make it configureable  
- [x] Implement expire date  
- [ ] Change Key to uuid
- [ ] Document it  
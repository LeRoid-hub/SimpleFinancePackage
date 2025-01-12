# SimpleFinancePackage


An Package to interact with an mock stock exchange

## Structure
| Field | Length |
| ----- | ------ |
| Key   | 8 Bytes |
| Operation | Byte |
| Asset | Byte |
| Price | Float64 |
| Volume | Float64 |
| CRC | 2 Bytes |

## TODOS
- [ ] Test it  
- [ ] Make it configureable  
- [ ] Implement expire date  
- [ ] Document it  
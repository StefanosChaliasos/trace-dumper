
object "Contract1" {
   code {
      datacopy(0, dataoffset("runtime"), datasize("runtime"))
      return(0, datasize("runtime"))
   }
   object "runtime" {
      code {
         switch shr(0xe0,calldataload(0)) 
            case 0xf8a8fd6d {
                let res0 := add(87, 21)
                mstore(0, res0)
            }
      }
   }
}


pragma solidity ^0.7.4;

library IcteLib {
    function RnTknGPlgn(bytes memory d) public returns (bool _s, bytes memory _r) {}
    function GtIeoBlnc(address clientAddress) public returns(int64) {}
    function Xfr(bytes memory fromAddrKey, address toAddress, int64 qty) public returns (bytes memory) {}
    function StSts(bytes memory k, byte s) public returns (byte) {}
    function RplcVl(bytes memory k, bytes memory v) public returns (bytes memory) {}
    function GtVl(bytes memory k) public returns (bytes memory) {}
    function RmVl(bytes memory k) public returns (bytes memory) {}
    function DbgLg(string memory l) public {}
    function PrgStrg(string memory k) public {}
    function StStrg(bytes memory ts, string memory k) public {}
    function GtStrg(string memory k) public returns(bytes memory) {}
    function StB(int32 i, byte v, string memory k) public {}
    function StS(int32 i, int16 v, string memory k) public {}
    function GtS(int32 i, string memory k) public returns (int16) {}
    function GtL(int32 i, string memory k) public returns (int64) {}
    function GtAddr(int32 i, string memory k) public returns(address) {}
    function GtI(int32 i, string memory k) public returns(int32) {}
    function AdVl(bytes memory k, bytes memory v) public returns (bytes memory) {}
}

contract LimitOrderBook {
    
    function onMessage(bytes memory goMsg) external {
        uint8 msgType = uint8(goMsg[0]);
        uint8 msgSubType = uint8(goMsg[1]);

        if(msgType == uint8(16) && msgSubType == uint8(0)) { //MsgFromClient            
            uint8 iMsgType = uint8(goMsg[22]);  // internal msgType 
            uint8 iMsgSubType = uint8(goMsg[23]); // internal msgSubType
            if(iMsgType == uint8(140)){ //MsgTypeTrade
                string memory tmpKey = "tmp";
                IcteLib.StStrg(goMsg, tmpKey);
                int64 cId = IcteLib.GtL(6, tmpKey);
                IcteLib.PrgStrg(tmpKey);
                bool succ;


                if (iMsgSubType == uint8(2)) { // NewOrder
                    IcteLib.RnTknGPlgn(goMsg);

                } else if (iMsgSubType == uint8(3)) { // CancelOrder

                } else if (iMsgSubType == uint8(4)) { // ReplaceOrder

                } else if (iMsgSubType == uint8(5)) { // AddToMarket

                } else if (iMsgSubType == uint8(6)) { // RemoveFromMarket

                } else if (iMsgSubType == uint8(7)) { // CancelAllOrders

                }


        }else{ // MsgFromGo 

        // pickup orderID and match to clientOrderID, we need to get coinbase from clientOrderID
            if (msgType == uint8(1) && msgSubType == uint8(1)) { // MsgTypeGetOrderID
                //Save orderID and clientOrderID in a map
                string memory tmpKey = "tmp";
                IcteLib.StStrg(goMsg, tmpKey);
                int32 cId = IcteLib.GtI(14, tmpKey); // clientOrderID
                int32 oId = IcteLib.GtI(18, tmpKey); // orderID
                IcteLib.AdVl(abi.encodePacked(oId), abi.encodePacked(cId));
                IcteLib.PrgStrg(tmpKey);
                bool succ;



            }else if(msgType == uint8(140)){ //MsgTypeTrade
                string memory tmpKey = "tmp";
                IcteLib.StStrg(goMsg, tmpKey);
                int64 cId = IcteLib.GtL(6, tmpKey);
                IcteLib.PrgStrg(tmpKey);
                bool succ;

                if (msgSubType == uint8(21)) { // OrderAccepted
                    IcteLib.DbgLg("orderAccepted");
                } else if (msgSubType == uint8(22)) { // OrderRejected

                } else if (msgSubType == uint8(23)) { // OrderOnMarket

                } else if (msgSubType == uint8(24)) { // OrderOffMarket

                } else if (msgSubType == uint8(25)) { // OrderCancelRejected

                } else if (msgSubType == uint8(26)) { // OrderCancelled

                } else if (msgSubType == uint8(27)) { // OrderReplaceRejected

                } else if (msgSubType == uint8(28)) { // OrderReplaced

                } else if (msgSubType == uint8(29)) { // OrderFilled

                }
            } else if(msgType == uint8(151)) { // MsgTypeQuote
                if (msgSubType >= uint8(0) && msgSubType <= uint8(14)){
                    

                }
            }
        }
    }
    }
}

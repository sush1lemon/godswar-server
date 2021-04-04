package opcodes

const (
	// ÍøÂçÁ¬½Ó¹¦ÄÜÏûÏ¢
	MSG_INVALID = 0

	//client<->login server
	MSG_LOGIN              = 1
	MSG_INVALID_CREDENTIAL = 3
	MSG_SELECT_SERVER      = 4 // ,...sush 1lemon..
	MSG_LOGIN_RETURN_INFO  = 6 // \...sush 1lemon.. | 0x5c 00 06
	MSG_REQUEST_GAMESERVER     //ÇëÇóÓÎÏ··þÎñÆ÷

	//loginserver<--->gamereserser
	MSG_VALIDATE_GAMESERVER = 300 //ÓÎÏ··þÎñÆ÷ÑéÖ¤Âë

	/*
		unknown opcode when entering the game
		10311, 10007, 12000, 10202, 10357, 10312

		10194 = walk?
	*/

	//client<->game server
	MSG_LOGIN_GAMESERVER    = 10000 //µÇÂ¼ÓÎÏ··þÎñÆ÷ / kick
	MSG_RESPONSE_GAMESERVER         //·µ»ØÑ¡ÔñÓÎÏ··þÎñÆ÷IP
	MSG_ROLE_INFO                   //½ÇÉ«ÐÅÏ¢
	MSG_CREATE_ROLE         = 10003 //´´½¨½ÇÉ«
	MSG_DELETE_ROLE                 //É¾³ý½ÇÉ«
	MSG_GAMESERVER_READY            //ÓÎÏ··þÎñÆ÷¾ÍÐ÷
	MSG_ENTER_GAME          = 10006 //¿Í»§¶Ë¾ÍÐ÷×¼±¸½øÈëÓÎÏ·
	MSG_CLIENT_READY                //¿Í»§¶Ë³õÊ¼»¯Íê±Ï
	MSG_GAMESERVER_INFO

	// ÐÅÏ¢¹¦ÄÜÏûÏ¢
	MSG_SELFINFO //×Ô¼ºµÄÐÅÏ¢
	MSG_OBJECTINFO
	MSG_LEAVE
	MSG_COLONY_LEAVE //ÈºÌåÏûÍö

	// ÓÎÏ·¹¦ÄÜÏûÏ¢
	MSG_WALK_BEGIN
	MSG_WALK_END
	MSG_SCENE_CHANGE

	// Õ½¶·¹¦ÄÜÏûÏ¢
	MSG_FIGHT
	MSG_ATTACK
	MSG_DEAD
	MSG_BACKHOME
	MSG_DROPS
	MSG_UPGRADE

	// ×°±¸µÀ¾ß¹¦ÄÜÏûÏ¢
	MSG_KITBAG
	MSG_STORAGE

	// ½»»¥¹¦ÄÜÏûÏ¢
	MSG_TALK = 10035
	MSG_TALKCHANNEL
	//MSG_TRADE

	// ·þÎñ¶Ë²ÎÊý£¬ÏµÊý
	MSG_PARAMATER
	// ¼¼ÄÜ¹¦ÄÜÏûÏ¢
	MSG_SKILL
	MSG_ACTIVESKILL_INFO
	MSG_PASSIVESKILL_INFO
	MSG_SELFPROPERTY
	MSG_EFFECT
	MSG_MAGIC_DAMAGE
	MSG_MAGIC_PERFORM
	MSG_MAGIC_CLUSTER_DAMAGE

	// Ö°Òµ¹¦ÄÜÏûÏ¢
	MSG_LEARN         //Ñ§Ï°¼¼ÄÜ
	MSG_SKILL_UPGRADE //¼¼ÄÜÉý¼¶

	MSG_PICKUPDROPS //Ê°È¡
	MSG_USEOREQUIP  //Ê¹ÓÃ»ò×°±¸
	MSG_MOVEITEM    //ÒÆ¶¯ÎïÆ·
	MSG_BREAK_ITEM  //²ð·ÖÎïÆ·
	MSG_STORAGEITEM //´æ´¢ÎïÆ·
	MSG_SELL        //ÂôÎïÆ·

	MSG_STALL        //°ÚÌ¯
	MSG_STALLADDITEM //Ìí¼ÓÎïÆ·
	MSG_STALLDELITEM //„h³ýÎïÆ·
	MSG_STALLITEM    //°ÚÌ¯ÎïÆ·
	MSG_STALLBUYITEM //Âò

	MSG_TALKNPC        //NPC¶Ô»°
	MSG_NPCDATA        //NPCÊý¾Ý
	MSG_SYS_NPC_DATA   //ÏµÍ³NPCÊý¾Ý
	MSG_SYS_FUN_USE    //ÏµÍ³¹¦ÄÜÊ¹ÓÃ
	MSG_NPCITEMDATA    //NPC··ÂôÊý¾Ý
	MSG_NPCSTORAGEDATA //NPC²Ö¿âÊý¾Ý
	MSG_NPCSELL        //NPC··Âô

	//ÈÎÎñ
	MSG_NPCQUEST             //ÈÎÎñ
	MSG_NPCNEXTQUEST         //ºóÐøÈÎÎñ
	MSG_NPCQUESTS            //ÈÎÎñÁÐ±í
	MSG_NPCQUESTSAVAILABLE   //ÈÎÎñË¢ÐÂ£¬¿É½Ó
	MSG_NPCQUESTSUNAVAILABLE //ÈÎÎñË¢ÐÂ£¬²»¿É½Ó
	MSG_NPCQUESTREWARD       //ÈÎÎñ±¨³ê
	MSG_NPCQUESTVIEW         //²é¿´ÈÎÎñÐÅÏ¢
	MSG_NPCACCEPTQUEST       //½ÓÊÜÈÎÎñ
	MSG_NPCCANCELQUEST       //È¡ÏûÈÎÎñ
	MSG_NPCCOMPLETEQUEST     //Íê³ÉÈÎÎñ
	MSG_NPCQUESTFAILD        //ÈÎÎñÊ§°Ü
	MSG_NPCREWARDQUEST       //Íê³ÉÈÎÎñ±¨³ê
	MSG_NPCQUESTKILLORCAST   //Í¬²½É±¹ÖÊÕ¼¯
	MSG_PLAYER_ACCEPTQUESTS  //Í¬²½½ÓÊÜÈÎÎñ
	MSG_FINDQUEST            //²éÕÒÈÎÎñ
	MSG_FINDQUESTRESULT      //²éÕÒÈÎÎñ½á¹û

	//HPMP»Ø¸´
	MSG_RESUNE

	//ºÃÓÑ
	MSG_RELATIONALL
	MSG_RELATION_REQUEST
	MSG_RELATION_RESPONSE
	MSG_RELATION_DELETE
	MSG_RELATION

	//½»Ò×
	MSG_TRADE
	MSG_TRADE_MONEY
	MSG_TRADE_ITEM
	MSG_TRADE_ADDITEM
	MSG_TRADE_REMOVEITEM

	MSG_EQUIPFORGE
	MSG_EQUIPFORGE_EQUIP
	MSG_EQUIPFORGE_MATERIAL
	MSG_EQUIPFORGE_EQUIPCANCEL
	MSG_EQUIPFORGE_MATERIALCANCEL
	MSG_EQUIPFORGE_CANCEL

	//±¦Ïä×ªÅÌ by lion
	MSG_GOLD_BOX

	MSG_EXPLORER_QUEST //Ì½Ë÷ÈÎÎñÑéÖ¤

	MSG_GOLD_BOX_RETURN

	MSG_TEAM_INVITE //client -> server  && server -> clientÑûÇë¼ÓÈë¶ÓÎé
	MSG_TEAM_REQUEST
	MSG_TEAM_INFO
	MSG_TEAM_ADD           //client -> server¼ÓÈë¶ÓÎé
	MSG_TEAM_DELETE        //client -> server¿ª³ý¶ÓÔ±
	MSG_TEAM_REPLACELEADER //client -> server¸ü»»¶Ó³¤
	MSG_TEAM_DISSOLVE      //client -> server¶ÓÎé½âÉ¢
	MSG_TEAM_LEAVE         //client -> serverÍÑÀë¶ÓÎé
	MSG_TEAM_TIP           //server -> client¶ÓÎéÌáÊ¾ÏûÏ¢
	MSG_TEAM_REJECT        //client -> server¾Ü¾ø¼ÓÈë¶ÓÎé
	MSG_TEAM_REFLASH       //server -> client¸üÐÂ¶ÓÎé
	MSG_TEAM_DESTROY       //server -> client¶ÓÎé½âÉ¢

	MSG_UPDATE_MP

	//¹«»á
	MSG_CONSORTIA_CREATE          //´´½¨¹«»á
	MSG_CONSORTIA_CREATE_RESPONSE //´´½¨ºó·µ»ØÖµ
	MSG_CONSORTIA_BASE_INFO       //¹«»á»á»ù±¾ÐÅÏ¢
	MSG_CONSORTIA_MEMBER_LIST     //³ÉÔ±ÁÐ±í
	MSG_CONSORTIA_INVITE          //ÑûÇë¼ÓÈë¹«»á
	MSG_CONSORTIA_DISMISS         //½âÉ¢¹«»á
	MSG_CONSORTIA_RESPONSE        //ÏìÓ¦ÑûÇë
	MSG_CONSORTIA_EXIT            //ÍË³ö¹«»á
	MSG_CONSORTIA_TEXT            //¹«»á¹«¸æ
	MSG_CONSORTIA_DUTY            //ÈÎÃüÖ°Îñ
	MSG_CONSORTIA_MEMBER_DEL      //ÒÆ³ý³ÉÔ±

	//¼ÀÌ³
	MSG_ALTAR_INFO //¼ÀÌ³ÏûÏ¢

	//·þÎñÆ÷´íÎó
	MSG_ERROR
	MSG_MANAGE_RETURN

	//¼¼ÄÜµãÊýÉý¼¶
	MSG_SKILLPOINT_UPGRADE

	//Í¬²½Êý¾Ý
	MSG_SYN_GAMEDATA

	//×´Ì¬
	MSG_STATUS

	//ÅÅ¶Ó
	MSG_LOGIN_QUEUE

	//·þÎñÆ÷Í¨Öª
	MSG_SERVER_NOTE
	MSG_SKILLBACKUP
	MSG_SKILL_INTERRUPT //¼¼ÄÜÖÐ¶Ï

	//Server->AS
	MSG_KEY_RETURN //GSÐ£ÑéÂë·µ»Ø
	MSG_BAN_PLAYER
	MSG_CONSORTIA_LVUP //¹«»áÉý¼¶
	MSG_ALTAR_CREATE   //´´½¨¼ÀÌ³
	MSG_ALTAR_LVUP     //¼ÀÌ³Éý¼¶
	MSG_ALTAR_OBLATION //¼ÀÌ³¹©·î

	MSG_MALLITEMDATA        //GameServer <---> client  ·µ»ØÓÎÏ·ÉÌ³ÇÎïÆ·ÁÐ±í
	MSG_ASSOCIATIONITEMDATA //GameServer <---> client  ·µ»Ø¹¤»áÉÌ³ÇÎïÆ·ÁÐ±í
	MSG_MALLSELL            //GameServer <---> client  ÓÎÏ·ÉÌ³Ç··Âô
	MSG_ASSOCIATIONSELL     //GameServer <---> client  ¹¤»áÉÌ³Ç··Âô
	MSG_ASSOCIATIONDISCOUNT //GameServer ----> client  ·þÎñÆ÷·¢¸ø¿Í»§¶ËµÄÉÌ³ÇÕÛ¿Û±í(Îª¶¯Ì¬¸Ä±äÕÛ¿Û)
	//ÉùÍû
	MSG_CRETIT_EXCHANGE     //ÉùÍû¶Ò»»
	MSG_QUESTEXPLORERRESULT //
	//Ôö¼ÓÒ»¸öÎïÆ·
	MSG_SYS_ADD_ITEM //Ôö¼ÓÎïÆ·
	MSG_SYS_DEL_ITEM //¼õÉÙÎïÆ·
	MSG_COUNT
	MSG_TARGETINFO
	MSG_DELAY_EXIT //µ¹¼ÆÊ±ÍË³ö

	MSG_WALK
)

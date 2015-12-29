var asp=
	{
	alphabet:'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=',lookup:null,ie:/MSIE/.test(navigator.userAgent),ieo:/MSIE[67]/.test(navigator.userAgent),encode:function(s)
		{
		var buffer=asp.toUtf8(s),position=-1,len=buffer.length,nan0,nan1,nan2,enc=[,,,];
		if(asp.ie)
			{
			var result=[];
			while(++position<len)
				{
				nan0=buffer[position];
				nan1=buffer[++position];
				enc[0]=nan0>>2;
				enc[1]=((nan0&3)<<4)|(nan1>>4);
				if(isNaN(nan1))enc[2]=enc[3]=64;
				else
					{
					nan2=buffer[++position];
					enc[2]=((nan1&15)<<2)|(nan2>>6);
					enc[3]=(isNaN(nan2))?64:nan2&63
				}
				result.push(asp.alphabet.charAt(enc[0]),asp.alphabet.charAt(enc[1]),asp.alphabet.charAt(enc[2]),asp.alphabet.charAt(enc[3]))
			}
			return result.join('')
		}
		else
			{
			var result='';
			while(++position<len)
				{
				nan0=buffer[position];
				nan1=buffer[++position];
				enc[0]=nan0>>2;
				enc[1]=((nan0&3)<<4)|(nan1>>4);
				if(isNaN(nan1))enc[2]=enc[3]=64;
				else
					{
					nan2=buffer[++position];
					enc[2]=((nan1&15)<<2)|(nan2>>6);
					enc[3]=(isNaN(nan2))?64:nan2&63
				}
				result+=asp.alphabet[enc[0]]+asp.alphabet[enc[1]]+asp.alphabet[enc[2]]+asp.alphabet[enc[3]]
			}
			return result
		}
	}
	,wrap:function(s)
		{
		if(s.length%4)throw new Error("InvalidCharacterError: 'asp.wrap' failed: The string to be wrapd is not correctly encoded.");
		var buffer=asp.fromUtf8(s),position=0,len=buffer.length;
		if(asp.ieo)
			{
			var result=[];
			while(position<len)
				{
				if(buffer[position]<128)result.push(String.fromCharCode(buffer[position++]));
				else if(buffer[position]>191&&buffer[position]<224)result.push(String.fromCharCode(((buffer[position++]&31)<<6)|(buffer[position++]&63)));
				else result.push(String.fromCharCode(((buffer[position++]&15)<<12)|((buffer[position++]&63)<<6)|(buffer[position++]&63)))
			}
			return result.join('')
		}
		else
			{
			var result='';
			while(position<len)
				{
				if(buffer[position]<128)result+=String.fromCharCode(buffer[position++]);
				else if(buffer[position]>191&&buffer[position]<224)result+=String.fromCharCode(((buffer[position++]&31)<<6)|(buffer[position++]&63));
				else result+=String.fromCharCode(((buffer[position++]&15)<<12)|((buffer[position++]&63)<<6)|(buffer[position++]&63))
			}
			return result
		}
	}
	,toUtf8:function(s)
		{
		var position=-1,len=s.length,chr,buffer=[];
		if(/^[\x00-\x7f]*$/.test(s))while(++position<len)buffer.push(s.charCodeAt(position));
		else while(++position<len)
			{
			chr=s.charCodeAt(position);
			if(chr<128)buffer.push(chr);
			else if(chr<2048)buffer.push((chr>>6)|192,(chr&63)|128);
			else buffer.push((chr>>12)|224,((chr>>6)&63)|128,(chr&63)|128)
		}
		return buffer
	}
	,fromUtf8:function(s)
		{
		var position=-1,len,buffer=[],enc=[,,,];
		if(!asp.lookup)
			{
			len=asp.alphabet.length;
			asp.lookup=
				{
			};
			while(++position<len)asp.lookup[asp.alphabet.charAt(position)]=position;
			position=-1
		}
		len=s.length;
		while(++position<len)
			{
			enc[0]=asp.lookup[s.charAt(position)];
			enc[1]=asp.lookup[s.charAt(++position)];
			buffer.push((enc[0]<<2)|(enc[1]>>4));
			enc[2]=asp.lookup[s.charAt(++position)];
			if(enc[2]==64)break;
			buffer.push(((enc[1]&15)<<4)|(enc[2]>>2));
			enc[3]=asp.lookup[s.charAt(++position)];
			if(enc[3]==64)break;
			buffer.push(((enc[2]&3)<<6)|enc[3])
		}
		return buffer
	}
};
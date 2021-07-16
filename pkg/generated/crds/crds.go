/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4V\xcdn+7\x0f\xdd\xcfS\x10\xf7[d\x13\x8fo\xf0uQ\xcc\xee\xd6i\x81\xa0\u007fA\x12dStAK\xb4\xcdF#\xa9\"ǩ\xfb\xf4\x85\xa4\x19{\xe2\xf8\x16\xd94\x9b`H\x8a\xa4\xce9\xa4\xdc,\x16\x8b\x06#?S\x12\x0e\xbe\x03\x8cL\u007f)\xf9\xfc%\xed˷\xd2rX\xeeo\x9a\x17\xf6\xb6\x83\xd5 \x1a\xfa\a\x920$C\xb7\xb4a\xcf\xca\xc17=)ZT\xec\x1a\x00\xf4>(f\xb3\xe4O\x00\x13\xbc\xa6\xe0\x1c\xa5Ŗ|\xfb2\xaci=\xb0\xb3\x94J\xf2\xa9\xf4\xfes\xfbM\xfb\xb9\x010\x89\xca\xf1'\xeeI\x14\xfb\u0601\x1f\x9ck\x00<\xf6\xd4\xc1\x1a\xcd\xcb\x10\x13\xc5 \xac!1I[M6\xf1\x9eRk\xbc\xd8\xd8\xee\xfbWLԚ\xd07\x12\xc9\xe4V\xb6)\fq:\xff\xb5\xe0Zel\xbd^\xfb\xbbr\xe0a*x(.Ǣ?^t\xffĢ%$\xba!\xa1\xbb\xd4pq\v\xfb\xed\xe00\xbd\v\xc8\x05ĄH\x1d\xac\xdc J\xa9\x01\x18q*\x8d-F$\xf67\xe8\xe2\x0eoj:\xb3\xa3\x1ek\xdf\x00!\x92\xffr\u007f\xf7\xfc\xff\xc77f\x00Kb\x12G-\x98_\xbdk\x1eX\x00\xc1Բ\x8b҅\x8542\xde\x02\xdci\x8e8Rja}\x00\xddј\an\v\xaa\x80>\x1f\xdaP\"oj\f<z\x8c\xb2\vz\r+\x17<\xfd\x90B?\x99J\xf8-9\xd2\\\xe1\xe9\x98\xed\x81\xe4\xd8\xd6\xd4B)\x8d\xec\xa5T5\x89,yet\xb0\t\tpD\x12NP\x8e\tO\x17\x1c;\xb4Y\xbcT\xb3T)\x80\xeePᕝ\x835\xc1 dA\x03(\xba\x97\xf2\u007fG\xb3\xac\x00\xbfzw8݉Դ\xb0z\x10ؤ\xd0W\x01E4%=*`\xa2\"\x17\xb2\xc0\x1e\xbe8\x17^\xc9\xfer\n\x9aj\xa2\xc9!\xc1_\x03oJ\xc1c\xa2\x8c9\xf8\xa0\x97\xcf\xe7\xd0\x10)\x95\xa9\xa9\xd96Ȯ\xbd:\x92\x1eS\xf6+O\xba\xae\u007fx\x9ei\xee\x04`\xa5\xfe\xcc\x04\xa0\x87,K\xd1\xc4~ۼw`Jx\x98\x978\xed\x967\xd1oU\x98\x85Z\xa3\xde\xf02J\x9e\xec\xa8m\b\x19\x16\x96LD\"!_\xd7L6\xa3\x87\xb0\xfe\x83\x8c\xb6\xf0H)\x1f\x04م\xc1٬\x97=%\x85D&l=\xff}\xcc&\x13\xad\x0e\x95$C\xab\x94<:أ\x1b躈\xb2Ǭ\xbc\x9c\x17\x06?\xcbPB\xa4\x85\x9fC\"`\xbf\t\x1d\xecT\xa3t\xcb\xe5\x96uڛ&\xf4\xfd\xe0Y\x0f\xcb2/\xbc\x1e4$YZړ[\no\x17\x98̎\x95\x8c\x0e\x89\x96\x18yQ\x9a\xf5ew\xb6\xbd\xfd\xdf$z\xb9\xba\x00\xf5;\x0e\xd6g\xa3\xbcr\xc8}\xf7\x91\x93e\xcd\xfd\v?y\xcf\xd5\xc5P\x8f\xd6\xfb\x9fhȦ\x8c\xe4\xc3\xf7\x8fO\xa7I-TUVN\xa1r\"(\x83\xcb~S&\x8f\xc7\xd1\xc9Y\xc8\xdb\x18\xd8k\x9dp\xc7\xe4\x15dX\xf7\xac\x99\xf9?\a\x12\xcdܵ\xb0*OM\x99\xd6hQɶp\xe7a\x85=\xb9\x15\n\xfd\xe7\xf4d4e\x91\xc1\xfb\x18A\xf3W\xf2<\xb8\xe24s\xe4e\x13+\x91\xf7\x98\xb0'\xa5t6\x8dhmy~\xd1\xdd_\x9c\xef\x0f\xcc\xebŲ\xf3E\xf9!\xfd\xc8ޜ\xbf#y\xa9|\xe0l\xe6\x93\x13\xcdԷx\xbf\x96f\xbe\x8b\x1a\x9f\xf9/\xa2v\xe6\x9f_\xaf\xf9*\x1c\x92Uj;\xd04P5hH\xb8\xa5\xd1\"\x8a:\x14\xb4\xd1\x18\x8a:\xf6;\xff\xd9\xf0\xe9ӛ_\x01\xe5\xd3\x04_9\x93\x0e~\xfb\xbd\xa9Y\xc9>O\x8f{6\xfe\x13\x00\x00\xff\xff\xb3\x96\xf8\xea\x95\t\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4U?s#\xb7\x0f\xed\xf7S`\xeeW\\sZ\x9d\xe7\x97\"\xb3\x9d\xa3Kq\x93?\xe39{\xdcdR@$$!\xe6\x92\f\x80\x95\xe2|\xfa\fII\x96e)\xe3&\xdb-\x00\x02\x0fx\x8f`7\x9b\xcd:\xcc\xfcH\xa2\x9c\xe2\x00\x98\x99\xfe2\x8a\xe5O\xfb\xa7\xef\xb5\xe74\xdf\xdetO\x1c\xfd\x00\x8bI-\x8d\xdfH\xd3$\x8e\xbeЊ#\x1b\xa7؍d\xe8\xd1p\xe8\x000\xc6dX\xccZ~\x01\\\x8a&)\x04\x92ٚb\xff4-i9q\xf0$5\xf9\xa1\xf4\xf6s\xff]\xff\xb9\x03pB\xf5\xf8\x03\x8f\xa4\x86c\x1e N!t\x00\x11G\x1a`\x89\xeei\xcaB9)[\x92g\x17\x90G\xed\x9b\xd9\voIz\x17\xd5\xe7~;\xeeP\xa8wi\xec4\x93+p֒\xa6|\xc8q-\xb8U\xda\xc3o\xad\xffP\x0f|;\x16]\x94\xa2\xd5\x1fX\xed\xa7\xeb1?\xb3Z\x8d\xcba\x12\f\xd7\xe0\xd7\x10帞\x02ʕ\xa0\x0e@]\xca4\xc0\xaf\x05^FG\xbe\x03\xd8O\xb0\u009d\xedg\xb4\xbd\xc1\x907x\xd3Һ\r\x8dغ\x01H\x99\xe2\xed\xdd\xd7\xc7\xff߿2\x03xR'\x9c\xad\xb2\xf1\xf1r7\xc0\n\x93\x92\aK\xe0\v\xff4G\xe7H\x15\xf0́\x1e\xe0\x16\"\xed\xde8`\xc7!\xc0\x92\x1a\xd3\xe4\x01vl\x1b\xb0\r\xc1KЗ\xca\xcd'X\by\x8a\xc6\x18\x00\xa3\x87\xdb\x10Ҏ\xfcq\x00ڒ\x11ۆ\xa4\xe4,Y\xe2\xc1\v\xb6A\xab\xa6s\f\xf7\x99\x1c\xc0\x0e\xf5\b\x82#$\xa9\xb1ok\x14\xf5\xf0\x8a[Եt=\xc0\xc3\x05\x17\xac\x98\x82o0\v\xc0)\xfbZ\xef\xd8sA\viu1\xef\x01]\xff\xf1HS\x96\x94I\x8c\x0f\xfal\x1f\x9ec>u\x02\xb0\xd1xf\x02\xb0\xe7\xa2%5\xe1\xb8\xee\xde:P\x04\x9fOK\xbc\xec\x89WѯuS\xa4բ\xf6\x02\xd1\xda\xd8^\xa4\xe4\xf7jl\r\xb3\x82P\x16R\x8ame\x143FH\xcb?\xc8Y\x0f\xf7$\xe5 \xe8&M\xc1\x97M\xb2%1\x10ri\x1d\xf9\xefc6-z,e\x02\x1a\xa9\x01G#\x89\x18`\x8ba\xa2OU9#>\x83P\xc9\vS<\xc9PC\xb4\x87_\x92\x10p\\\xa5\x016fY\x87\xf9|\xcdv\u0601.\x8d\xe3\x14ٞ\xe7u\x9d\xf1r\xb2$:\xf7\xb4\xa50W^\xcfP܆\x8d\x9cMBs\xcc<\xab`c݃\xfd\xe8\xff'\xfb\xad\xa9\x1f/\x8c\xfa\r\a\xcb3\x1d\f\xef9T7տPS\xb6T\xb9\xc0\xb8?\xdaZ\u007fa\xa0\x98\xea%\xfc\xf1\xfe\x01\x0ex\x1bK\x8d\x90\x97P}\xe1\xa6̕㊤E\xae$\x8d5\vE\x9f\x13\xc7v\xfd\\`\x8a\x06:-G\xb6B\xfa\x9f\x13\xa9\x15\xdazX\xd4\x17\xe3\xe4j\xf4\xf05\xc2\x02G\n\vT\xfaϙ)\xd3\xd4Y\x19\xde\xfb\xb89}\xec\u0383ۜN\x1ce\x85\xe7\xc6\xe1\x1d\n\x8ed$g\x17\x11\xbd\xaf\xaf(\x86\xbb\x8bW\xfb\x1dW\xf5b\xd9\xd3=\xfa\x0e\xfd\x14NX\xe8DA\xb3\xcb\xe8\xcf\xfc\xa7e\xba\xab\xb0\xb4\xa8\xc5\x0f`2Q3X\x12\\\xd3ޢ\x866ծ\xcbs\x92m\xbf\xc9N_\xe0\x0f\x1f^=\xa3\xf5ץ\xd8f\xa7\x03\xfc\xf6{ײ\x92\u007f<\xbc\x88\xc5\xf8O\x00\x00\x00\xff\xff(|\x8f\x15\xe4\b\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4XM\x93\xdb6\x0f\xbe\xfbW`\xf2\x1e\xf6\xedL,'\xd3\x1e:\xbee\xbcM\xebi\x93\xeed3{\xe9\xf4@\x91\xb0ŮD\xb2\x04\xe4\xd4\xed\xf4\xbfw@J\xb6,\xd9\xebl?|\x13\t\x02\xe0\x03\xe0\x01\xe8\xd9|>\x9f\xa9`\x1f0\x92\xf5n\t*X\xfc\x8d\xd1\xc9\x17\x15\x8f_Sa\xfdb\xf7z\xf6h\x9dYª%\xf6\xcd\a$\xdfF\x8d\xb7\xb8\xb1β\xf5n\xd6 +\xa3X-g\x00\xca9\xcfJ\x96I>\x01\xb4w\x1c}]c\x9co\xd1\x15\x8fm\x89ekk\x831)\xefM\xef^\x15_\x15\xaff\x00:b:\xfe\xd16H\xac\x9a\xb0\x04\xd7\xd6\xf5\f\xc0\xa9\x06\x97\xa0k\xefp\x13}CN\x05\xaa<SQ*\xfd\xd8\x06\x13\xed\x0ec\xa1\x1d\x99P\xec\x9aO*b\xa1}3\xa3\x80Z\\\xd9F߆%<-\x9c\xadt\xaew\xd7\x16\x83o\xa3o\xee;\x83i\xaf\xb6\xc4ߟ\xdf\xff\xc1R\x96\tu\x1bU}\xce\xe5\xb4M\xd6m\xdbZ\xc53\x023\x00\xd2>\xe0\x12ދ;Ai43\x80\x0e\xad\xe4\u07bc\xc3c\xf7ZաR\xaf\xb3J]a\xa3\xb2\xf7\x00>\xa0{s\xb7~\xf8\xf2\xfed\x19\xc0 \xe9h\x03'\xe4o\xa67\x00K\xd0\x12\x1a`\x9fだ\xc0\xe1'\x88]\xf0\xe1\xff\xbc\x0fV\xab\xbaރ\x82\xbb\x87\xd5\x17 \ue0c2\xfe\x06\x05\xc0\x8fN#p\x85Ы\xbd\xb9!\xb8\xab\x14!T\x8a\x00\x1a\xbf\xcb&\xfa}F\x036\x19ߩ\xda>a=\xd9\x12ͽ5X\xdf\xde\x1cn\x17\xa2\x0f\x18\xd9\xf6a̿A\x9e\x0fV\xc7X\b\\Y\n\x8c$8R\xb2\xd3\x01\x8f\xa6C\x18\xfc\x06\xb8\xb2\x04\x11CDB\x97S^\x96\x95\x03_\xfe\x82\x9a\v\xb8\xc7(\a\x81*\xdf\xd6F*a\x87\x91!\xa2\xf6[g\u007f?h#\xb9\xa9\x98\xa9\x15#1X\xc7\x18\x9d\xaa\x05\x86\x16_\x82r\x06\x1a\xb5\x87\x88\xa2\x17Z7АD\xa8\x80w>\"X\xb7\xf1K\xa8\x98\x03-\x17\x8b\xad徆\xb5o\x9a\xd6Y\xde/R9ڲe\x1fiap\x87\xf5\x82\xecv\xae\xa2\xae,\xa3\xe66\xe2B\x05;OκT\xc7Ec\xfe\xd7CO7'\xe0\xf1^r\x948Z\xb7\x1dl\xa4\xc2y\x02e)\x1c\t\xb3\xea\x8e\xe6[\x1c\xc1\x94%\xc1\xe3\xc37\xf7\x1f\x8fQO\x80gl\x8f\xa2t\x84Y \xb2n\x831K\x1e\x92\x04\x9d\t\xde:N\x1f\xba\xb6\xe8\x18\xa8-\x1b\xcb\x12\xbf_[$\x96\b\x14\xb0J\xe4\x05%B\x1b\x8cb4\x05\xac\x1d\xacT\x83\xf5J\x11\xfe\xe7 \v\x9a4\x17\xf0>\x0f\xe6!\uf3853N\x83\x8d\x9e\x06/\xc4d\xc2\x01\xf7\x01u:d7\x16\xe9\x98֒\xab%f\xca2㪇\xf5m\x01\xf0\xb1Bx\xd7\xf9\x96\xa2R\"\xf8\x1d\xc6h\x8dA\xf72\xc5a\xe3c\xa38\xd7\x11\x1en\x02\xc7\bw\xa6u\x01\xf0\xe6n\xfd\xad\x90w*\x84\x94;ys\x9f\xce\xca}E\xcfѽL\x1a\xc5\xc9eϓBG\fI\xfbx}\x04\xd0\xc1\x89\xce\xe5CZ\x96(\xe9\x9am\x9a\x89\x8e\v\xa1\x93_\xeeD\x1f0x\xb2\xec\xe3\xfe\x8a\x03\x82j>\"\xd9ߝ\x91\xebF\xe4hq\x87\xa7\x8c(\x91\xe9b\xe1\xfa\x1er\xc2Ƌ\xbb\x87\x15\xd4v\x87\x04\xd6A\xd3\x12C\xa5v\bJk\xa4\x03%\x1dM=\xe7j);V\xcai\xac\xafܪ\xf7&\v\x83u\xc6ja\xc1\xbe2S<\xf3\x9ew[/P\x0fZ\xcc\xe8\xb4VNr\x8d\x90A1(\xb7g\xdbH\x806R\xb7'\xe8DT\xba\x92\xb4\x06\xc6\xd8X!\xdbP\xa5\x1a\x87\xf5\xe6TTzU\x167\x13\xf1\v\x98\x94\xdeר\xdchwʊ\x134zb\x1c\xe6\xf5?O\xb4\xf3<!\xbf\\\x85K(\xf7|\xe9.g5\xf6\xe0\xaco\xa7:/\x1e\x93\x88ڈ#\f\xe6\x87\x02\x1c-\x8f\xcbc\xb4=H\xb1ю\xe0<Z:\xfa\xfbYTɊ\xdb\x11O\\\xe6\x0f\xed\x9bP\xe3\xe9\xc8\xfat\x94W\xd3\x13i(\x88\xa6\x8b\xbcd\xadr\xc7\x14\xfc\xa4\xa87#MIj\x9aRӻ\xa1,\xdc\xcfk\x1b\x1f\xcfi\xa7\x8b\x81\x9767\x17\x15\x13\t\x19\xbaUY\xe3\x128\xb6\xcfJ\x8e\x06\x89\xd4\x16\xaf\x80\xf0.K\xe51\xa0;\x02\xaa\xf4mߧ\xbd\xc3\x1b\xeabQ<\xc7~\xaa\xcc+\xd6\xf3\x18\xdaU\x9ancLC\x01ˬ\xd9\xf1\xfb\xa4!>ˇ\xbeh\xbfS\xce\xd4ל\x91xVIp\xd2\\\xb8R]\xf8\a\xfdvHO\x13͗\x13\x15\x9elv\x13\xaf\x0e\r\xafC)\xbd\x9fR\x86\x9di\u007f\x117\x18\xd1\xe943mN\xce\xcaH\xd5\xcf\x11&7\xff\xc3g&\xbc\xd4{J\x99\xacr8\x84\xad\xdfܭ\xb3\xc5\x02\xde\xfa(\\\x0e\x9e\xab<\xdbE3\x0f*\xf2>\x05\x80^\x9eX\xebIf\x1a\xad+\x11\x83\x8b\f=A\xe63X\xfa\x88\xc7\xdf\xf1#\xbd\xec\xae\xfb!o\xc3\xde\x0f9\xf2/\xfbq\x9e\xaf\xe1<\xc7\xe6e\xf1\xe2B\x99LX\xf6\xc2\xc6\xd4\xea<\xe5\xcb\xec\xe2\xa9D\x85f\xc0S\xc4>\n\x01\rV\xda\xf2\xf0\x82\xe9\x15w,\x0f\u007f\xfc9;\x12\xbe\f?\x81Ѽ\x1f\xff\v\xf0\xe2\xc5ɓ>}j\xef\x8c\xcd\u007fu\xc0O?ϲa4\x0f\xfd+]\x16\xff\n\x00\x00\xff\xff\xc0ψ[d\x11\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4V=o\xe46\x13\xee\xf5+\x06\xf7\x16n^i\xcfH\x8a@]\xb0Na$w0l\xc3M\x90\x82Kή\x18S$33\xd4\xc5\t\xf2\xdf\x03\x92\x92\xf7\xcb\xf6\xf9\x8a\xa8\x13\xe7S\xcf3\xf3PM۶\x8d\x8a\xf6\x01\x89m\xf0=\xa8h\xf1OA\x9f߸{\xfc\x81;\x1bV\xd3e\xf3h\xbd\xe9a\x9dX\xc2x\x8b\x1c\x12i\xbc\u00ad\xf5Vl\xf0͈\xa2\x8c\x12\xd57\x00\xca\xfb *\x1fs~\x05\xd0\xc1\v\x05\xe7\x90\xda\x1d\xfa\xee1mp\x93\xac3H%\xf9Rz\xfa\xd8}\xdf}l\x004a\t\xbf\xb7#\xb2\xa81\xf6\xe0\x93s\r\x80W#\xf6`С {\x15y\b\xc2\xddF\xe9\xc7\x14\r\xd9\t\xa9ӞM\xec\xa6\xf1\x8b\"\xect\x18\x1b\x8e\xa8s\x1f;\n)\xf6\xf0\xb6s-1\xf7]\xbf\xf9\xaaT\xbb\x9b\xab\x15\x83\xb3,?\xbf`\xfc\xc5ru\x88.\x91rg\x9d\x16\x1b[\xbfKNѩ\xb5\x01`\x1d\"\xf6\xf09\xb7\x10\x95F\xd3\x00\xcc\xf0\x94\x96\xda\x19\x80\xe9R\xb98\xa8˚O\x0f8\xaa\xda1@\x88\xe8\u007f\xbc\xb9~\xf8\xee\xee\xe8\x18 R\x88Hb\x97\x8f\xab\xcf\x01\xf5\a\xa7\x00\x06Y\x93\x8dR\x88\xb9\xc8\t\xab\x17\x98\xcc92ȀKkh\xe6\x1e lA\x06\xcb@\x18\t\x19}\x9d\x82|\xac<\x84\xcd逸\x83;\xa4\x1c\b<\x84\xe4L\x1e\x8e\tI\x80P\x87\x9d\xb7\u007f=gc\x90P\xca8%\xc8\x02\xd6\v\x92W\x0e&\xe5\x12\xfe\x1f\x9470\xaa' \xccy!\xf9\x83\fŅ;\xf8\x14\b\xc1\xfam\xe8a\x10\x89ܯV;+\xcbX\xeb0\x8e\xc9[yZ\x95\t\xb5\x9b$\x81xepB\xb7b\xbbk\x15\xe9\xc1\njI\x84+\x15m[\x9a\xf5e\xb4\xbb\xd1\xfc\x8f\xe6E\xe0\x8b#\xf0\xe4)\xb3\xc8B\xd6\xef\x0e\fe\x9c\xde@9O\x14X\x065\x87֯\u0603\x99\x8f2\x1e\xb7?\xdd\xdd\xc3R\xba\x02^\xb1ݻ\xf2\x1e\xe6\f\x91\xf5[\xa4깥0\x96,\xe8M\f\xd6Ky\xd1\u03a2\x17\xe0\xb4\x19\xadd\xfe\xfeHȒ\x19\xe8`]\xf6\x196\b)\x1a%h:\xb8\xf6\xb0V#\xba\xb5b\xfc\xcfA\xcehr\x9b\xc1{\x1ḟRt\xea\\q:0,\xe2\xf0\n'w\x11u\xa6\xa4`T\xb4o\x0f|\x0e=\x8a|y\xc3\xf2S5\xe7\x16c`+\x81\x9eN\xed'U\xef\a\x9cC2\xa3sL\xde\x06B!\x8b\x13\x96~\x16\xdd(\x94v%\xc8/\xc2Q\x1c\x16UZ\xdd<\xac\xc1\xd9\t\x19\xac\x871\xb1\xc0\xa0&\x04\xa55\xf2\xf3\x96\xed+\x9d5\xf7\n\xd0\x05\xbf\xb9\xc6\xf5\xd5\xf97\xbd\x1a\x96\xc7\xcb\x12\x9e,C{\x06Ӊy_\xeb]̊\x92\xc4op\xbbNDe\xe8\x8bc\xd5.\x9c5y\x0fn\xa6\xb4\xa8\xd8;\xa9\xd6a\x8c\x0e\x8fo\xaf\xb7\xd9^\x9fG\x141$S\xe7N\xec\x88Y?\x8fo\v\xf8\xa2x)\x96W2\xd3\xcfe\xe5/\xb8\x86X\x86\xc4h`\x1b\xe8\xa5\x1a|\xd6\xd56Ш\xa4\x87\xbc\xe4mNq\xe6\x91oa\xb5q\u0603P:7\xbf1'#2\xab\x1d~\x05\x8aOի\x8a\xe0\x1c\x02j\x13\x92\xbc\xc4\xcd\x05\xcf\xdcu\xdf\xd2I\x1c\x14\u007f\xad\x8f\x9b\xec\xb3\xdf\xfb\xfd\x9c\xe02&\xf5\xde\u007f^\xb1o\xe8\xe0\xc5q=߈\xf6X_\u03a2\n\xd7\xe6\x80\b\x96@\x19დ\xb4y\xbe\xa0\x96\xc4\xf3V\xc0\xdf\xff4\xfb\x05\xc9B\x10\x05\xcd\xe7\xd3_\x9f\x0f\x1f\x8e\xfefʫ\x0e\xde\xd8\xfas\a\xbf\xfe\xd6\xd4\xc2h\x1e\x96ߔ|\xf8o\x00\x00\x00\xff\xff\x86~\xe9IV\n\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xb4XOo\xe3\xc6\x0e\xbf\xfbS\x10y\x87\\by\x17o\xf1\xf0\xe0\xdb\xd6\xe9\x9f`\xbb\x8b \tr)z\x18\xcd\xd0\xd64\xa3\x19uH9\xeb\x16\xfd\xee\x05g$\xcb\xf2\xbf$E\xd77\x8dH\x0e\xf9#\xf9#\xe5\xc9t:\x9d\xa8\xc6>b$\x1b\xfc\x1cTc\xf1+\xa3\x97'*\x9e\xfeO\x85\r\xb3\xf5\xfbɓ\xf5f\x0e\x8b\x968\xd4wH\xa1\x8d\x1a\xafqi\xbde\x1b\xfc\xa4FVF\xb1\x9aO\x00\x94\xf7\x81\x95\x1c\x93<\x02\xe8\xe09\x06\xe70NW苧\xb6Ĳ\xb5\xce`L\xc6\xfb\xab\xd7\xef\x8a\x0fŻ\t\x80\x8e\x98\xd4\x1fl\x8dĪn\xe6\xe0[\xe7&\x00^\xd58\a\xf2\xaa\xa1*0\x15\xa5\xd2Omc\xa2]c,\xb4'\xd3\x14\xeb\xfaYE,t\xa8'Ԡ\x16\x0fV1\xb4\xcd\x1c\xce\vg\xe3\x9d\xc79\xda\xfb\xee\x9et\xe4,\xf1\xa7\xd1\xf1ϖ\xf2\xabƵQ\xb9\x1d\xbf\xd2)Y\xbfj\x9d\x8a\xc3\xf9\x04\x80thp\x0e_\xe4\xaaFi4\x13\x80\x0e\x80t\xf5\xb4\vq\xfd^\xb9\xa6R\xef\xb3%]a\xad\xb2g\x00\xa1A\xff\xf1\xf6\xe6\xf1\xbf\xf7\xa3c\x00\x83\xa4\xa3m8\x81y\xb9u\x13,AKh\x80\x03D\xfc\xbdEb\xe0J1\xa8\xadc\"\xc2\xea\t}\x01p\x93\x9e|\xe0\xadR\xad\xbcZ!p\x85`\xfd\x1a=\x87\xb8\x81\xb0\x1c\xc2\x05\xe5\r\x98\x80Y\r<f=\xfcj\x89\xc1z\b\xd1`\x94\x13\xed\x82φbWB\xb0\x8c\xa1\xde\xf1\xe4r\x1bM\x13C\x83\x91m\x9f\x92\xfc\xdb)՝\xd3\xfd\xd8\x05\x9e,\x05Fj\x14)]\xda\x01\x8d\xa6CT\x82\xe0\xca\x12Dl\"\x12\xfa\\\xb5r\xac<\x84\xf27\xd4\\\xc0=FQ\x04\xaaB\xeb\x8c\x14\xf3\x1a#CD\x1dV\xde\xfe\xb1\xb5F\x12\xa1\\\xe3\x14c\n\x9c1z\xe5`\xad\\\x8bW\t\xa4Zm \xa2\u0605\xd6\xefXH\"T\xc0\xe7\x10\x05\xe4e\x98C\xc5\xdc\xd0|6[Y\xee\xdbP\x87\xban\xbd\xe5\xcd,u\x94-[\x0e\x91f\x06\xd7\xe8fdWS\x15ue\x195\xb7\x11g\xaa\xb1\xd3\xe4\xacO\xadX\xd4\xe6?=\xeat9\x02\x8f7R\x93\xc4\xd1\xfa\xd5\u038b\xd4\x04gP\x96n\x90ZQ\x9dj\x8eb\x00S\x8e\x04\x8f\xbb\xef\xef\x1f\x86\x84'\xc03\xb6\x83(\r0\vD\xd6/\xa5`D2ՇXAo\x9a`=\xa7\a\xed,z\x06j\xcb\xda2\xf5e-\x19(`\x91\xf8\aJ\x84\xb61\x8a\xd1\x14p\xe3a\xa1jt\vE\xf8\xcdA\x164i*\xe0\xbd\x0e\xe6]\xea\xdc\x17\xce8\xed\xbc\xe8)\xedDN\xee\x1bԩ\x99\x05\xa3\xc4\xd5\x03\xf0\xa2:\xd2<\xdea\xf2\xcbLy\x87M +ݾ\xff~\xefև\n;\x15\xc9h\xa7#\xdd0\xf0\x8b\x97\xcc$A\xdfS_r\xb2'\xa9\xd9\xed\xe3\x02\x9c]#\ta\xd4-1Tj\x8d\xa0\xb4F\xdav\xd6`\xfd\xc0\xa1\x13\xe0ʯG\xe0'\xe5\x8d\xc3\x17b\xb9\x1b\tCĥ\x14%\aP\xf0\xa9-1zd\xa4\xad\xc9+\xd0m\x8c\xe8\xd9m@\x81\xc4P\xb6\x9c\vW\xea\xb8DH\xc3Ԡ\x91\xb0$\x84e+us\xe0\xc3\xe9\\@f\xbd\x1f\xd3\f;\xf2n\xcf\xff\x8f\xb77I\xb4\xaf\x824\xfb`\x19\xe2\x98vK\x94\xeeLѡשG\x96#]i!\xa9\x18\xbb\xb4h\xae\x92\xf2\xf6\x11R\xe7\xa7$\x95\xd8\a\xa6\xa5\xab>\xde\xde\xe4\x1b\v\xf8!DP~\x03\x81\xab\xdc\xcb\xd1L\x1b\x15y\x93REW\xa3ۤ\x81mDS\x1c\r\xf0Ln\xe1(O\x1dE\xa6\xa7+qV,\n͟\xc4\xe3\x9f\xf8\x91&\xf7\xcb~\xc8\xec\xef\xfd\x10\x95\u007fُ\x1e\xcaCO\xa6\t\xa9#\xc7\xe2ŉ\x86:  \xf9\xf5m\xbdP^\xa3{\xa1\xa1\xeeG\xc2`\xbd\xb1Z\xe5\x1e\xeav\x91\x00:\xbf\v~\x15$\xfc\xde~\x01\xfb\xdaZy)9B\x06Y`\xfc\x86m-\xa0-\xa5\xf8R\x8d\xf6\x8c\x13Q\xe9\ne81\xc6\xda\xca\x1cn\xaaD\xffR\xe9#\xd1JQ'n\x0e\xc4O\xc0R\x86\xe0P\xf9\xc9\xcb\xc0O\x0f\xc8t\xef\xf5\x98\x9d^5\x03XqKg\xa6\xc0\"sR'\x98\xb7\x9c\x9dx\x85iҢ\xf3\xcai\xa0C\xdd8\x1c/\xe4\xe7s\xbe8\xd4H\xfbR4]\xfbI֔\x1f\\zV\xd4_#\\$s\x82\xd2>pIY\xb8\xdfa\x85ȎX\xa7\x03\u007f\x96!֊\xe7 \x1b\xc0TL\x1cH\xc8'\x85*\x1d\u0381c{*\xd1G\x9b\xacF\"\xb5zi\x92|\xceRyC\xeaT@\x95\xa1\xe5Q:.\xa9\xcb\xd3!\xf7\x9du\xe1\xd8\xd2pć,\xb6\x1d\x00\xfd\xad\x8c\xe6xo\x0fȕ\x1b~\x13*\xa9_^\xf0\xe7Vd\x86\x05e(S쫴o\xf87\xa1\xd1İ\x8aHGf\xe7\xf8\xf6Nl\x8bF۸\xa0\x0e\x19\xf1\xfc0\x16\\\xe8:\xf8\x13\\\xdf\xc3g=\xff\xef\xc3\x19\xfa\x96σ\x15\xc6#\x12\x1cX\xb9\xef\xe4\x96os\xc3+\x98\xfd\xe6\xfa\x95\xac\x0e7\xd7\xf9cRH\xb4D\xf4\xdb\xef\xc8\aY\x82\x9e\xadsB\xd8K\xeb\\ށ\x9e+\xcc\xfbB*\x17X\xc9W#\a\xb8\xb8\x1f\n\xf3\xe2-\xa9\xa7\xb5\xeeU\xbf\x1c\x9d\xbf#\xb7e)Y\xa5\xa9\xa3]K\x8c\x91\xae\x80d\xcf\xdc\x1d\xc6\xdd6\x13\x91\x9a\xe0M\x1aHm\x83qm)\xc4^o\xa0.\xd1*r\xb4;\x1f\xda\x1c\x95~ک\xb1\x9e\x8b\xb7\x9f3\x87&\xdfP\xf1G3x8\u007f\xa6\xe3\x9d\xff@+Q\xac\xd9\xe1?\xe2\x10\x85\xd8vN\xdar\xfb\xd1\xd8\x1b\xee\xe6\x0f\xfc\xf9\xd7d\x18E\xb2\xa87\x8c\xe6\xcb\xfe\x9f(\x17\x17\xa3\xffHң\x16\\\xf3\x1fD\xf0˯\x93|1\x9a\xc7\xfe\x8f\x109\xfc;\x00\x00\xff\xff\xbe\x95\x0f\xa8\x9a\x12\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdo\xe3\xba\x11\xbf\xfb\xaf\x18\xbc\x1er\xa9\xe5\xb7}\x8b\xa2\xf0\xad\x9b\xec\x02A\xbbA\x10os)z\xa0ȱņ\"\xb5\x9c\xa1\xbd\xee__\f)\xd9\xf2G\xe2\xb8\xe8>_\x82P\xf3\xc5\xdf\xfc\xe6C\x9aL\xa7Ӊ\xea\xec3F\xb2\xc1\xcfAu\x16\u007f0z\xf9\x8f\xaa\x97\xbfPe\xc3l\xfda\xf2b\xbd\x99\xc3m\"\x0e\xed\x13RHQ\xe3\x1d.\xad\xb7l\x83\x9f\xb4\xc8\xca(V\xf3\t\x80\xf2>\xb0\x92c\x92\u007f\x01t\xf0\x1c\x83s\x18\xa7+\xf4\xd5K\xaa\xb1N\xd6\x19\x8c\xd9\xf8\xe0z\xfdk\xf5\xb1\xfau\x02\xa0#f\xf5o\xb6Eb\xd5vs\xf0ɹ\t\x80W-\xce\xc1\x84\x8dwA\x19\xaa\xc4e\x1b\xd6\x18+\xed\xc9tպݨ\x88\x95\x0e\xed\x84:\xd4\xe2~\x15C\xea\xe6\xf0\x86d1\xdb\xc7Z\xeey\xd7{\xc8G\xce\x12\xff\xed\xe0\xf8\xef\x968?\xea\\\x8aʍ\"ʧd\xfd*9\x15\xf7\xe7\x13\x00ҡ\xc39<\x88\xabNi\x94\xb3\xfe\xea\xd9\xf5\xb4\xbf\xdc\xfa\x83r]\xa3>\x14K\xba\xc1V\x95\xc8\x00B\x87\xfe\xaf\x8f\xf7Ͽ-\x0e\x8e\x01\f\x92\x8e\xb6\xe3\f\xe3\x10d\u007fZ#(X\xa3\xc3\x18\xa6\x9dK+\xeb!\"q\x88\xb8S\xefb\xe80\xb2\x1d0(\xbf\x11+F\xa7G\xcen$\x9e\"\x05F\xe8\x80\x04\xdc\xe0p34\xfd\x15 ,\x81\x1bK\x10\xb1\x8bH\xe8\vA\xe4Xy\b\xf5\xbfQs\x05\v\x8c\xa2\bԄ\xe4\x8c\xf0f\x8d\x91!\xa2\x0e+o\xff\xb3\xb3F\xc0!\xbbq\x8a\x91\x18\xacg\x8c^9X+\x97\xf0\x8f\xa0\xbc\x81Vm!\xa2\u0605\xe4G\x16\xb2\bU\xf05D\x04\xeb\x97a\x0e\rsG\xf3\xd9ley`\xbc\x0em\x9b\xbc\xe5\xed,\x93\xd7։C\xa4\x99\xc15\xba\x19\xd9\xd5TE\xddXF\xcd)\xe2Luv\x9a\x83\xf5\x99\xf5Uk\xfe\x10\xfb\x1a\xa1\x9b\x03\xf0x+$ \x8e֯F\x0f2\xeb\xde@Y\xe8\a\x96@\xf5\xaa\xe5\x16{0\xe5H\xf0x\xfa\xbc\xf8\x06\x83\xeb\x02x\xc1v/J{\x98\x05\"\xeb\x97\x18\x8b\xe42\x866[Ao\xba`=\xe7\u007f\xb4\xb3\xe8\x19(խe\xc9\xdf\xf7\x84Ē\x81\nns\xa9C\x8d\x90:\xa3\x18M\x05\xf7\x1enU\x8b\xeeV\x11\xfet\x90\x05M\x9a\nx\xef\x83yܥ\x8e\x85\vN\xa3\aC\x03y%'\x8b\x0e\xb5\xa4$c\x94\xdb\xe2\x1exQ=\xd0<_a\xf2\xab\x95~I\xdd\x13v\x81,\x87\xb8\x95\xfep,s\xe4\xf9ӑ\x8aX_[\x83\xd4\x1b\x93\\\x0f\x8f\x04{X\x86\xb8\xebDՉ\xbax\x1c.\"-\xa8\x94*\x9e\xc8U'Q\xbd\x82\xb2\xfc\xb4\v\x1e\x85P\v\xaf:j\x02?\xe1\x12#z}\xe9r\xb7\xa2\xf8\xe5\x9c\xe28\xc6\xdcAs\x99\xef\x1cQ/_\xee WΤ\x1e\xee=\x10\xb7\x82oM~\xdc*\x16\x8b'\xfev\rzv\xf6\x11\xdcg\xb5Dh\xa4\t\x15\xde\xe7\xb8v\x9e\x88\x15'\x02\xeb\xfb\xfa9\n\xf0*\x1c\xbb\x18\x84\xffh>{\xb6\xbc\xbd\xbf\xbb\x00\xdf\xe3\xb1\xfc\x80\x9a5R7K\x8b\xb1\xc7\x06\xf7\xb6\x01\xb3\xb0@cI\x14<\xa2)\xf7\x93\x89\xb9\x89\x96\x05l\xc0\x1f\x96r\xabY\a\x97Z\xbc\xea\"\xfd\xccُ\xf4\xb7\xef\xf1t$\x9eg@4\xe5.l\xdb\x02yo\x146\x8a@+\xe7\xa4\x01Iz)7\xb8\x1b*\x92C\xba\xe4\xdeC*w\x86O\xc2(\xd4\xc8\v\x03NE\xff\x9a[\x0e9\xbe\x98\xa7\xc5N\xf0\x8d\x04\xed(\xddW\xe4հ\xbf\xd6\xd92A\xdf\xe8m\xc3\n\xb1\xe8\x99<t\xb9\x18\xf3((\xa72\xbaw\x92\xd5;\u06dd\x0em\xe7\xf0p\xb9\xbb\xd0\x10N5Nɠ\xfc\xbe\xfc2\x19\x8a\x92\xf0a\xaf\xbfcCQ\x17ޯ\xd1C\xf0\xb0T֡\x19\xed\x95\x17Xt&&\xfa\x1f\x88$K\xad\xaa\x1d\u0381c\xba\x8ag-\x12\xa9ե^\xfa\xb5H\x95šW\x01U\x87\xc4\a\r\xeb\x86\xfa\x94^U\xd0\x1e\u007f\xf0\x13rܾ7\x91\x0f'\nÞW\xe3.\x93\xe5\x9c\x1bi\xce\xdeX=4W\xf1\x06Q\xb4s\x02\x8e\x12\x06\xb7O\x15\xfc\xa3o\xc9K\xeb\x18#\x1c\xdfr\x18\x00\xb0i\xacn\x84!\x98{t\x8dKi!#\a\x12\xc7)\x14?3\x9d]\xa3\xe8R2\x1fE\xe6\\)\xee\x06\xf6\xf9Z\x94\x1f\xfaԞ\x9a\x9f\xc2\x03nΜ\xde\xfb\xc7\x18V\x11\xe9\x94\xd3Ӂ\xfah\xce<\xcb\xd9=s\xfe%g\xebʱ\xa7\x91\xe4=\xea!\x98K\xc8H\xad\xde)V_\x95W+\x8c\xe0\x83\xc1B\xa1F\x11tV\xbf\xa0\x81\xd4\x1d`\x94Y\xb4\xf7ҏ\xbe\x8dun\xb4͂\"\xa0\x10\xbc\xfc=P\xb6c\xb3ǖ\xee\xfbl\x8c\"\xd2Bs\u007fÃ\xdca\x18\x14Z\x19e\x8a\x82\a˻ \xf6\x1e\xea-(\x1f\xb8\xe9\xefv\xed\x02\x91Syyo\xc8b\xd0\x0474\xd6\xc0ʁOm-մ\x84z\xcbH\x87\xa3(o`c>\ue947ќ\xc3g\xa4\x1ea\xad|\x06\xb8\xafUc\xa9sj\xbb\x8b2\xef\xabRgҮ\xf7\x9dj0&S%?;\x85\xe0\xf5\x91#\xbf\x1c\xc2]\xf0g\xa8\x04\xa3ڶ\x9e\xff\xfc\xf1\xacDAX\xde3W\x18\xcfHd\xb0>\x89\x97\x9f\xe3\xe1\x95I\x0ey\xb3⸽\r\xc9\xf3ŝj\x10<\x18\xa0\xe3\x9c\r\xed\x8f21#N\x951\xb2\xf0\xa9\xa1\xe3\xde\xed;n\xdfH\x87V\x9c\x841\xe0\x917!\xbe\x80%J\x98\xdfG\xe4\xf4{\u0084}\x87\x16É\xe4\x9d3*\xfd\x92\x8d{\x03\x06\xeb\xb4ZI\xfd\xbc\xdax\xad\xe7\xdf\xfe\xf4\n0\xe7a#V\x91\xdf;\x9e\x16\a\u0097W\x8cl\xfc\x1d\v\xe7\x81\xd9\xdfwK(Ezq\x17}\xee\xc5\xde\xd8D\xfb\x024\xff\xa7\x15TF\xb1\x8d8\xfa\xe41=|g>\xd1\xca\x00\x9b\x11\x04\x12\x8fl@\xe5d\xbf\xd1*\xad\xb1c4\x0fǟ\xf4~\xf9\xe5\xe0\x8b]\xfeW\aol\xf9P\t\xff\xfcפXE\xf3<|\x96\x93\xc3\xff\x06\x00\x00\xff\xffS\x8a'\x11\"\x15\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcX\xcdr#\xb9\r\xbe\xeb)P\x9b\x83/Q{'\xbb\x95J\xe9\x96\xd5l\xaa\\\xc98S\xf6\xcc\\R9\xb0I\xb4\x9a1\x9b\xe4\x12\xa4l\xe5\xe9S ٭֏-9\x95I_l\x81\x00\b~\x00>\xa2{\xb1\\.\x17\xc2\xebo\x18H;\xbb\x02\xe15\xbeD\xb4\xfc\x8b\x9a\xa7?Q\xa3\xdd\xed\xf6\xc3\xe2I[\xb5\x82u\xa2\xe8\x86\a$\x97\x82ď\xd8i\xab\xa3vv1`\x14JD\xb1Z\x00\bk]\x14,&\xfe\t \x9d\x8d\xc1\x19\x83a\xb9A\xdb<\xa5\x16ۤ\x8d\u0090\x9d\x8f[o\u007fl~n~\\\x00Ȁ\xd9\xfc\x8b\x1e\x90\xa2\x18\xfc\nl2f\x01`ŀ+H\xde8\xa1\xa8\xe1\r\a\xb7\xc5\xd0HK\xca7\xdb\xe1Y\x04l\xa4\x1b\x16\xe4Q\xf2\xe6\x9b\xe0\x92_\xc1\x1b\x9a\xc5i\x8d\xb4\x9c\xf2k\xf6\x9f\x05FS\xfc\xebL\xf87M1/x\x93\x820S,YF\xdan\x92\x11a\x94.\x00H:\x8f+\xb8\xe7-\xbc\x90Ȳz\xe0\xbc\xe5\xb2\x1ei\xfbA\x18ߋ\x0fŏ\xecq\x10%\"\x00\xe7\xd1\xfe\xf9\xf3ݷ\x9f\x1e\x0f\xc4\x00\nI\x06\xedc\x06\xaf\x84We-\x82\x80-\x1a\fn\xe9M\xdah\v\xad\x90O\xc9O\xb6>8\x8f!\xea\xf1\xe0\xe5\x99\x15\xc2Lz\xb4\xd3\r\aS\xb4@q\x05 A\xecq<\x16\xaa\x1a?\xb8\x0eb\xaf\t\x02\xfa\x80\x84\xb6\xd4\x04\x8b\x85\x05\xd7\xfe\vel\xe0\x11\x03\x1b\x02\xf5.\x19ť\xb2\xc5\x10!\xa0t\x1b\xab\xff=y#\x88.ocDD\x8a\xa0m\xc4`\x85\x81\xad0\t\u007f\x0f\xc2*\x18\xc4\x0e\x02\xb2_Hv\xe6!\xabP\x03\x9f\\@жs+\xe8c\xf4\xb4\xba\xbd\xdd\xe88\x16\xb9tÐ\xac\x8e\xbb\xdb\\\xaf\xbaM\xd1\x05\xbaU\xb8EsKz\xb3\x14A\xf6:\xa2\x8c)\xe0\xad\xf0z\x99\x83\xb5\xb9ЛA\xfd.Զ\xa0\x9b\x03\xf0\xe2\x8e+\x80b\xd0v3[ȥ\xf6\x06\xca\\u\xa0\tD5-\xa7\u0603\xc9\"\xc6\xe3\xe1\xd7\xc7/0n]\x00/\xd8\xeeUi\x0f3C\xa4m\x87\xa1hv\xc1\r\xd9\vZ坶1\xff\x90F\xa3\x8d@\xa9\x1dt\xe4\xfc\xfd\x96\x90\"g\xa0\x81u\xeenh\x11\x92W\"\xa2j\xe0\xce\xc2Z\fhւ\xf0\xbb\x83\xcchҒ\xc1\xbb\x0e\xe691\x1d+\x17\x9cf\v#k\xbc\x92\x93G\x8f\x92S\x921\xcaL\xb8\a\x9eM\x0f,\xcfw\x18?\xa5\x15\x1f\xd0;\xd2х\xdd\xf1\xfaѮ\xbf\x1c\xa9\xb3\xe7\xadVH\xd5\x11\xe7y\\bܡs\xa1\x12P\x03_\tU\x16\f\xc9D\xed\r\x9e\x1a5'ۿ\x02\xe5>\xf6=1_\x13\xfa\xa4\x9d\xbb:\xa8\x02`\xd4\x03\xe6\u007fj@ς@\nc\xb8\xa2\xbe\xf4\b\x94+\xf6\x86\x8a\xa2&H\xe3Q\x1e\xad\xf0Ի8\xf9=\t\xa2sa\x101\xd3>.\xd9\xfe=G\xa4\xea\xfe\xee\xe3\x85\xd3=N\x8acQhŕ\xdai\f9P\x16\x8d\xde\n\x1b\"l\x9dI\x03\xbe\v\xf3\xd1\xc5\x03v\x18\xd0J\xbc2\xaeI\u007f\fώ\xb7Pf\xcb)2\x16\xd7x\x19\xe6r\x91Ԧ/\xb9(x\xb2\x9b\xd1\xf7t\xa1ݎ\x92\xf5\x03\xcb\xe0.N\xb9\x8a\xaerD\u07bc\xfa\xa5(b\"\xd0\xf6\x00\x9dw\xe1Q<\xad\x85\x95h.@\xf1u\xa6\n\xda*-\xf9\xf6\x18\x0f\xc7\x11ʲ\xe6\xec\xc61\xa3־y%\x9c\xd69\x83\xc2^E%\xf9\x9co\x90I\t\xed\xb1\xc21\x92J\b\x99y\x8b\x94oʪ\xd7\\\xc9-\xd2\r\xde\xe0\xe1\xf0\xf46D\xebS\x8b\xd3>\x15v\xcc_n\xd3b\u009d\xba\xb7\x9e\xfa\xb4\x18\xa3\x02ܢ\x05g\xa1\x13ڠ\x9a\xa6\xb6\v\xdd}&\x1e\xfa/\x1a\x9c\aF\xd1\x1a\\A\f\xe9]\xfd_\x93\xc0\xdc\xf5\xf7\xae\xbb\x84ށ\xf2\x01p\xcck\xae\xeb\x18\x81\x801\xec\xf2\xe9\x0e\xb0h\xe0!/\xb8\t\xdd:\x05\xb9\x16w\x80/\xdeY\xe6\x13a&_\x03\xca^XM\xc3i\x85\x8e\x90h\x1b\u007f\xfa\xc3+\xe7\xe5\xa1i\x83\xe1hu@\"\xb1\xb9\xc4*\x9f\x8aV\x19I\xaa\t\x88֥8k\xef\x1b\xaa\xb5\xfb\xae\x86\xb6\xf8\x123\x12\xd7\xd6\xec\xfd\x89\xc1\x88\\\x8bS\xd1\x16y외j\xe7\x17\x16ė\xf8ZF`\xfdP\xef\xcc\xe8\xa0\xd3&b\x80\xc3\x13N\xdc\xf1\xdck\xd9s+`f\xb3\x16;\x1e|f\xce9\x86\xd7\xd3\xf4=*\xd7\xf7\x82.\xa5\xf13\xeb\x9c\xe3\x1b\x1co\xa8s\x84\xc3\x0f\xda4\x9c:_\xc2=>\x9f\x91\xde\xd9\xcf\xc1m\x02\xd2i\xf3.\xc7\x1eGuf\xadl\xffk\b\xee\xb8N\xb3%\xf3o\xf2\u007f\xc9I;\xb7\x9e\xe9\xfc\x8d\xa5S\xd4ކ48\x89\xc4/v\xf7N]\u0096\x89\xed\xa3\x88⓰b\x83\x01\xacSX\n\xb0\x17\x04^˧\\h3\x94s\x05\xee\xf7`n\xd4\x04\xcfژل\r\x82\x80\x9c\xb3\xfcwf\xaa\xe7.\x8f\xfdܕ\\Σ\x91\xdc \xf6&\x8ez\xf3\x10\xc8\r\xfcj!\xc8Y\xd0q\n`\xef\xbf݁\xb0.\xf6\xf5T\xefjo_\v\xe1ReV5\xe8\x9d\x19\xef\x1e\x17\x85\x01\x9b\x86\x96\xfb\xb0\x83vǷ\xf7\xc1$\x95G\x99y-ϴ\xc5>\xfa\x88T\xc1\x95\xc2flk\x93+Mވ\xdd\x14d\x1e\xa1\xb9E\xf9J\xdb\xd3[\x1dR\xf9\xda\xcdK\xa7\x00\xbc~'\xf3\x93#\xf8\xe8\xec\x99\x12\x82C\xf2\xfe\xe3\xcfg5\xde\"\xf0\xbc\xceP\xfd»|\x9f\x1d^\x19s\xf8\xc9t\xb7v\xc9\xc6\v\x19~\x98\x14\x0f.\xca}\xc6\xf6\xc4I\xb9,\x85R<\x94\x89\x91\xa4k\xc5\x16\xea\xad2\x95\xb8N\xc0b|v\xe1\t4Q*\xe9b\xe9o\t\x13\xceނ\x12\xf1\x9bo\x10\xf2)\xfb\xb5\n\x14\xb6i\xb3\xe1\x9e\xf9\x1fި\x14E\x88\xd7^f\x8f\aʗf\xaf\xec\xfa\x8aw\xa4\x03\xa7\xff\xcf\x01\xeal\xa1\xf0\xa5\xa9\x03\xce>x,\x0fߘO\xac\xf2\xd1\xd4ls\x8a.\xf0\x94R$\xfb\xf1ZH\x89>\xa2\xba?\xfe\x8a\xf7\xc3\x0f\a\x1f\xea\xf2O\xe9\xac\xd2\xe5\xcb$\xfc㟋\xe2\x15շ\xf1\x8b\x1c\v\xff\x13\x00\x00\xff\xff\x81b\x1b\xc9\x13\x15\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1.CustomResourceDefinition))
	}
	return objs
}

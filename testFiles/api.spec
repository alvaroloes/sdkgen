GET https://www.alvarloes.com/posts
<- [
	{
		"id":"1234",
		"author":{
			"name":"John",
			"age":20
		},
		"title":"We really need a client SDK generator",
		"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
		"comments":[
			{
				"id":"4567",
				"title":"I like it",
				"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
			},{
				"id":"4567",
				"title":"I like it too",
				"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
			}
		]
	},{
        "id":"12345",
        "author":{
            "name":"John",
            "age":20
        },
        "title":"We really need a client SDK generator",
        "body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
        "comments":[
            {
                "id":"4567",
                "title":"I like it",
                "body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
            },{
                "id":"4567",
                "title":"I like it too",
                "body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
            }
        ]
    }
]

GET https://www.alvarloes.com/posts/:id
<- {
	"id":"1234",
	"author":{
		"name":"John",
		"age":20
	},
	"title":"We really need a client SDK generator",
	"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
	"comments":[
		{
			"id":"4567",
			"title":"I like it",
			"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
		},{
			"id":"4567",
			"title":"I like it too",
			"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
		}
	]
}

POST https://www.alvarloes.com/posts
-> {
	"title":"We really need a client SDK generator",
	"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us"
}

<- {
   	"id":"1234",
   	"author":{
   		"name":"John",
   		"age":20
   	},
   	"title":"We really need a client SDK generator",
   	"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
   	"comments":[
   		{
   			"id":"4567",
   			"title":"I like it",
   			"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
   		},{
   			"id":"4567",
   			"title":"I like it too",
   			"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
   		}
   	]
}

PUT https://www.alvarloes.com/posts/:id
-> {
	"title":"We really need a client SDK generator. Please"
}

<- {
   	"id":"1234",
   	"author":{
   		"name":"John",
   		"age":20
   	},
   	"title":"We really need a client SDK generator. Please",
   	"body":"(...) we to make the machine work for us, thus we should write generators to make the computer write the non-creative part of the code for us",
   	"comments":[
   		{
   			"id":"4567",
   			"title":"I like it",
   			"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
   		},{
   			"id":"4567",
   			"title":"I like it too",
   			"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
   		}
   	]
}

DELETE https://www.alvarloes.com/posts/:id

GET https://www.alvarloes.com/posts/:id/comments
<- [
	{
		"id":"4567",
		"title":"I like it",
		"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
	},{
		"id":"4567",
		"title":"I like it too",
		"body":"I like it, but it seems to be really hard to come up with a powerfull and flexible generator"
	}
]

GET https://www.alvarloes.com/posts/:id/comments/:id
<- {
	"id":"4567",
	"title":"I like it",
	"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
}

POST https://www.alvarloes.com/posts/:id/comments
-> {
	"title":"I like it",
	"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
}
<- {
	"id":"4567",
	"title":"I like it",
	"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
}
PUT https://www.alvarloes.com/posts/:id/comments/:id
-> {
	"title":"I really like it"
}
<- {
	"id":"4567",
	"title":"I really like it",
	"body":"I like this post about api generators. It would be awesome to have a powerfull generator to avoid coding SDKs for all the client laguanges your API target"
}
DELETE https://www.alvarloes.com/posts/:id


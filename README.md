# journey
app repository


# api specification

---------------------
post api/v1/users
	params:
	{
		mobile: "",
		password: "",
	}
	response
	{
		state_code:
		message:
		data:{
			id: 
			username:
			mobile:
		}
	}

---------------------
post api/v1/sessions
	params:
	{
		field_value: username/mobile/email
		password:
	}
	response
	{
		state_code:
		message:
		data:{
			id: 
			username:
			mobile:
		}
	}

---------------------
get api/v1/apps?page_no={:no}&page_num={:num}
response body:
	{
		state_code: 
		message:
		data:{
			{
				id:
				logo_url:
				tags:[]
				dlink:
			}

		}
	}

---------------------
get api/v1/apps/:id
	{
		state_code: 
		message:
		data:{
			id:
			logo_url:
		}
	}

---------------------
get api/v1/categories
	{
		state_code: 
		message:
		data:{
			id:
			name:
		}
	}

<!-- get api/v1/categories/:id -->


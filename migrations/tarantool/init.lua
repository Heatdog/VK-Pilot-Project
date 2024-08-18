box.schema.space.create('users', {
    if_not_exists = true
})

box.schema.space.create('data', {
    if_not_exists = true
})

box.space.users:format({
    {name = 'id', type = 'string'},
    {name = 'login', type = 'string'},
    {name = 'password', type = 'string'}
})

box.space.data:format({
    {name = 'key', type = 'string'},
    {name = 'value', type = 'any'}
})

box.space.users:create_index('primary', {parts = {'id'}, if_not_exists = true})
box.space.users:create_index('login', {parts = {'login'}, if_not_exists = true})

box.space.data:create_index('primary', {parts = {'key'}, if_not_exists = true})
box.schema.space.create('users', {
    if_not_exists = true
})

box.space.users:format({
    {name = 'id', type = 'string'},
    {name = 'login', type = 'string'},
    {name = 'password', type = 'string'}
})

box.space.users:create_index('primary', {parts = {'id'}, if_not_exists = true})
box.space.users:create_index('login', {parts = {'login'}, if_not_exists = true})

box.space.users:insert{"1", "admin", "presale"}
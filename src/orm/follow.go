package orm

// IsFollow func
func IsFollow(ui, oi string) bool {
	var c int
	var f = new(Follows)
	Connect().Model(f).Where("sub_id = ?", ui).Where("obj_id = ?", oi).Count(&c).Close()
	return c > 0
}

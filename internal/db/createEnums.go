package db

import "gorm.io/gorm"

func CreateEnums(db *gorm.DB) {
	createEnum(db, "role_enum", "'admin', 'teacher', 'student', 'staff', 'parent'")
	createEnum(db, "parent_type_enum", "'father', 'mother', 'student', 'guardian', 'relative', 'other'")
	createEnum(db, "income_type_enum", "'business', 'job', 'agriculture', 'pension', 'foreigner', 'other'")
	createEnum(db, "grade_point_enum", "'gpa', 'cgpa', 'marks', 'percent'")
	createEnum(db, "status_enum", "'active', 'inactive', 'suspended'")
	createEnum(db, "gender_enum", "'male', 'female', 'other'")
	createEnum(db, "staff_position_enum", "'teacher', 'principle', 'accountant', 'librarian', 'receptionist', 'clerk', 'peon', 'driver', 'security', 'cleaner', 'cook', 'shopkeeper', 'wathman', 'electrician', 'sweeper', 'nurse', 'care_taker', 'lab_assistant', 'other'")
	createEnum(db, "qualification_type_enum", "'ssc', 'hsc', 'diploma', 'bachelor', 'master', 'phd', 'other'")
	createEnum(db, "discount_type_enum", "'flat', 'percentage'")
	createEnum(db, "payment_status_enum", "'pending', 'success', 'failed', 'refund', 'cancel', 'hold', 'charge'")
	createEnum(db, "payment_method_enum", "'cash', 'cheque', 'card', 'bank_transfer', 'mobile_banking', 'wallet', 'other'")
}

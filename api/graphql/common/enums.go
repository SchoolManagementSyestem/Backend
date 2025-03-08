package common

import "github.com/graphql-go/graphql"

// Gender Enum
var GenderEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "GenderEnum",
	Values: graphql.EnumValueConfigMap{
		"MALE":   &graphql.EnumValueConfig{Value: "male"},
		"FEMALE": &graphql.EnumValueConfig{Value: "female"},
		"OTHER":  &graphql.EnumValueConfig{Value: "other"},
	},
})

// Status Enum
var StatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "StatusEnum",
	Values: graphql.EnumValueConfigMap{
		"ACTIVE":    &graphql.EnumValueConfig{Value: "active"},
		"INACTIVE":  &graphql.EnumValueConfig{Value: "inactive"},
		"SUSPENDED": &graphql.EnumValueConfig{Value: "suspended"},
	},
})

// User Role Enum
var UserRoleEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "UserRoleEnum",
	Values: graphql.EnumValueConfigMap{
		"ADMIN":   &graphql.EnumValueConfig{Value: "admin"},
		"TEACHER": &graphql.EnumValueConfig{Value: "teacher"},
		"STUDENT": &graphql.EnumValueConfig{Value: "student"},
		"STAFF":   &graphql.EnumValueConfig{Value: "staff"},
		"PARENT":  &graphql.EnumValueConfig{Value: "parent"},
	},
})

// Staff Position Enum
var StaffPositionEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "StaffPositionEnum",
	Values: graphql.EnumValueConfigMap{
		"PRINCIPLE":     &graphql.EnumValueConfig{Value: "principle"},
		"ACCOUNTANT":    &graphql.EnumValueConfig{Value: "accountant"},
		"LIBRARIAN":     &graphql.EnumValueConfig{Value: "librarian"},
		"RECEPTIONIST":  &graphql.EnumValueConfig{Value: "receptionist"},
		"CLERK":         &graphql.EnumValueConfig{Value: "clerk"},
		"PEON":          &graphql.EnumValueConfig{Value: "peon"},
		"DRIVER":        &graphql.EnumValueConfig{Value: "driver"},
		"SECURITY":      &graphql.EnumValueConfig{Value: "security"},
		"CLEANER":       &graphql.EnumValueConfig{Value: "cleaner"},
		"COOK":          &graphql.EnumValueConfig{Value: "cook"},
		"SHOPKEEPER":    &graphql.EnumValueConfig{Value: "shopkeeper"},
		"WATCHMAN":      &graphql.EnumValueConfig{Value: "watchman"},
		"ELECTRICIAN":   &graphql.EnumValueConfig{Value: "electrician"},
		"NURSE":         &graphql.EnumValueConfig{Value: "nurse"},
		"CARE_TAKER":    &graphql.EnumValueConfig{Value: "care_taker"},
		"LAB_ASSISTANT": &graphql.EnumValueConfig{Value: "lab_assistant"},
		"SWEEPER":       &graphql.EnumValueConfig{Value: "sweeper"},
		"OTHER":         &graphql.EnumValueConfig{Value: "other"},
	},
})

// Grade Point Enum
var GradePointEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "GradePointEnum",
	Values: graphql.EnumValueConfigMap{
		"GPA":     &graphql.EnumValueConfig{Value: "gpa"},
		"CGPA":    &graphql.EnumValueConfig{Value: "cgpa"},
		"MARKS":   &graphql.EnumValueConfig{Value: "marks"},
		"PERCENT": &graphql.EnumValueConfig{Value: "percent"},
	},
})

// Discount Type Enum
var DiscountTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "DiscountTypeEnum",
	Values: graphql.EnumValueConfigMap{
		"FIXED":      &graphql.EnumValueConfig{Value: "fixed"},
		"PERCENTAGE": &graphql.EnumValueConfig{Value: "percentage"},
	},
})

// Parent Type Enum
var ParentTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "ParentTypeEnum",
	Values: graphql.EnumValueConfigMap{
		"FATHER":   &graphql.EnumValueConfig{Value: "father"},
		"MOTHER":   &graphql.EnumValueConfig{Value: "mother"},
		"GUARDIAN": &graphql.EnumValueConfig{Value: "guardian"},
		"RELATIVE": &graphql.EnumValueConfig{Value: "relative"},
		"OTHER":    &graphql.EnumValueConfig{Value: "other"},
	},
})

// Income Type Enum
var IncomeTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "IncomeTypeEnum",
	Values: graphql.EnumValueConfigMap{
		"BUSINESS":    &graphql.EnumValueConfig{Value: "business"},
		"JOB":         &graphql.EnumValueConfig{Value: "job"},
		"AGRICULTURE": &graphql.EnumValueConfig{Value: "agriculture"},
		"TUTION":      &graphql.EnumValueConfig{Value: "tuition"},
		"PENSION":     &graphql.EnumValueConfig{Value: "pension"},
		"FOREIGNER":   &graphql.EnumValueConfig{Value: "foreigner"},
		"OTHER":       &graphql.EnumValueConfig{Value: "other"},
	},
})

// Qualification Type Enum
var QualificationTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "QualificationTypeEnum",
	Values: graphql.EnumValueConfigMap{
		"JSC":      &graphql.EnumValueConfig{Value: "jsc"},
		"SSC":      &graphql.EnumValueConfig{Value: "ssc"},
		"HSC":      &graphql.EnumValueConfig{Value: "hsc"},
		"DIPLOMA":  &graphql.EnumValueConfig{Value: "diploma"},
		"BACHELOR": &graphql.EnumValueConfig{Value: "bachelor"},
		"MASTER":   &graphql.EnumValueConfig{Value: "master"},
		"PHD":      &graphql.EnumValueConfig{Value: "phd"},
		"OTHER":    &graphql.EnumValueConfig{Value: "other"},
	},
})

// Transaction Status Enum
var TransactionStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "TransactionStatusEnum",
	Values: graphql.EnumValueConfigMap{
		"PENDING": &graphql.EnumValueConfig{Value: "pending"},
		"SUCCESS": &graphql.EnumValueConfig{Value: "success"},
		"FAILED":  &graphql.EnumValueConfig{Value: "failed"},
		"REFUND":  &graphql.EnumValueConfig{Value: "refund"},
		"CANCEL":  &graphql.EnumValueConfig{Value: "cancel"},
		"HOLD":    &graphql.EnumValueConfig{Value: "hold"},
		"CHARGE":  &graphql.EnumValueConfig{Value: "charge"},
	},
})

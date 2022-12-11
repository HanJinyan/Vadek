package main

import (
	"Vadek/configuration"
	"Vadek/database"
	"Vadek/loger"
	"go.uber.org/fx"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {

	_ = fx.New(
		fx.Provide(configuration.NewConfiguration),
		fx.Provide(loger.NewLogger),
		fx.Provide(database.NewGormDB),
		fx.Provide(loger.NewGormLogger),
		fx.Populate(&Db),
	)

}
func grenerate() {
	g := gen.NewGenerator(gen.Config{
		Mode:              gen.WithDefaultQuery,
		OutPath:           "./database",
		ModelPkgPath:      "./model/entity",
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	g.UseDB(Db)
	g.ApplyBasic(
		g.GenerateModel("attachment", gen.FieldType("type", "consts.AttachmentType")),
		g.GenerateModel("category", gen.FieldType("type", "consts.CategoryType")),
		g.GenerateModel("comment", gen.FieldType("type", "consts.CommentType"), gen.FieldType("status", "consts.CommentStatus")),
		g.GenerateModel("comment_black"),
		g.GenerateModel("journal", gen.FieldType("type", "consts.JournalType")),
		g.GenerateModel("link"),
		g.GenerateModel("log", gen.FieldType("type", "consts.LogType")),
		g.GenerateModel("menu"),
		g.GenerateModelAs("meta", "Meta", gen.FieldType("type", "consts.MetaType")),
		g.GenerateModel("option", gen.FieldType("type", "consts.OptionType")),
		g.GenerateModel("photo"),
		g.GenerateModel("post", gen.FieldType("type", "consts.PostType"), gen.FieldType("status", "consts.PostStatus"), gen.FieldType("editor_type", "consts.EditorType")),
		g.GenerateModel("post_category"),
		g.GenerateModel("post_tag"),
		g.GenerateModel("tag"),
		g.GenerateModel("theme_setting"),
		g.GenerateModel("user", gen.FieldType("mfa_type", "consts.MFAType")),
	)
	g.Execute()

}

#### 模型定义
User 模型类的 BeforeCreate 方法可以在数据插入数据库之前,做一些操作.
```
    import (
        "github.com/jinzhu/gorm"
        "github.com/satori/go.uuid"
    )

   func (model *Model) BeforeCreate(scope *gorm.Scope) error {
        uuid := uuid.NewV4()
        return scope.SetColumn("Id", uuid.String())
   }
```

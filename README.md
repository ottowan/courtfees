# Function And Parameter Naming Conventions

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
datecalculate
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
datecalculate.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
datecalculate
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_ActionToTest_Input_152_Should_Be_150
```

## Variable Name
- ชื่อตัวแปรเป็นคำเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```

- ชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ **camelCase** เช่น
```
startDay, endMonth
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
days, months
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
ResultDay, ResultMonth
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน

## คำสั่ง Run Test
### ค่าสั่ง Run Acceptance Test (Robot)

```
robot duration.robot
```

### คำสั่ง Run Acceptance Test (API)
```
newman run filename
```

## Additional
[See more git and go command](https://github.com/ImKK-000/git-and-go-step)

## Ebook Thai
https://www.dropbox.com/sh/is3hwdqa1dpsb99/AACb9QvxPEUo1Z-tD0Nz2KNZa?dl=0

## GO Video Thai
[Basic]
https://www.youtube.com/playlist?list=PLEE74DyIkwEm4BHNx0_0gPb-TlHFH4ZSc

[Web]
https://www.youtube.com/playlist?list=PLEE74DyIkwEmCIvkQ7xpS9dZ7aSEO39hf

[Authenticate]
https://www.youtube.com/watch?v=XseIOD9SnfE&t=1s
https://www.youtube.com/watch?v=S_hSblZrcak
https://www.youtube.com/watch?v=XdF99kbWQ5Q

[GO+Bootstrap]
https://www.youtube.com/watch?v=tfNa0_i8BjM&t=535s

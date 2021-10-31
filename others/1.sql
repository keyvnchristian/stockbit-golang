SELECT u1.id, u1.UserName, u2.UserName as ParentUserName
FROM USER u1
LEFT JOIN USER u2
ON u1.ID <> u2.ID AND 
u1.Parent = u2.ID

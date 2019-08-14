export const Services ={"persons":{"general":{"name":"persons","description":"person service","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{}},"members":{"general":{"name":"members","description":"The members of a project","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{"List":{"data":{"request":null,"response":"person_collection"},"query":{"q":{"description":"Query term to search a member","type":"string","meta":{"label":"Search","hint":""},"__proto":{"type":"string"}}},"deeplink":{"rel":"list","href":"/api/members","method":"GET"}}}},"project":{"general":{"name":"project","description":"Project data","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{}},"projecttasks":{"general":{"name":"projecttasks","description":"The tasks of a given project","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{}},"projecttree":{"general":{"name":"projecttree","description":"The tree with relevant project data","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{}},"tasks":{"general":{"name":"tasks","description":"The tasks","version":"1.0.0","lifecycle":{"deprecated":false,"info":"This version is still valid"},"__proto":{"imports":[]}},"services":{}}}
export const Types ={"furo.link":{"name":"Link","type":"furo.link","mime":"application/furo.link+json","description":"link","fields":{"rel":{"description":"the relationship","type":"string","__proto":{"number":1}},"method":{"description":"method of curl","type":"string","__proto":{"number":2}},"href":{"description":"link","type":"string","__proto":{"number":3}},"type":{"description":"mime type","type":"string","__proto":{"number":4}}}},"furo.meta.field.constraint":{"name":"Meta","type":"furo.meta.field.constraint","mime":"application/furo.meta.field.constraint+json","description":"constrains of fields","fields":{"constraint":{"description":"constrain of a field","type":"keyValuePair","__proto":{"number":1,"map_from":"string","map_to":"string"}}}},"furo.meta.field.meta":{"name":"Meta Field","type":"furo.meta.field.meta","mime":"application/furo.meta.field.meta+json","description":"meta","fields":{"label":{"type":"string","description":"meta information of a field","__proto":{"number":1}}}},"furo.meta.field":{"name":"Meta","type":"furo.meta.field","mime":"application/furo.meta.field+json","description":"fields of meta info","fields":{"meta":{"description":"meta information of a field","type":"furo.meta.field.meta","__proto":{"number":3}},"constraints":{"description":"constraints of a field","type":"furo.meta.field.constraint","__proto":{"number":4,"repeated":true}}}},"furo.meta":{"name":"Meta Field","type":"furo.meta","mime":"application/furo.meta+json","description":"meta info","fields":{"meta":{"description":"fields of meta info","type":"keyValuePair","__proto":{"number":1,"map_from":"string","map_to":"furo.meta.field"}}}},"furo.reference":{"name":"Reference","type":"furo.reference","mime":"application/furo.reference+json","description":"reference","fields":{"display_name":{"description":"String representation of the reference","type":"string","meta":{"readonly":true},"constraints":{},"__proto":{"number":1}},"id":{"description":"Id of the reference","type":"string","__proto":{"number":2}},"rel":{"description":"the relationship","type":"string","__proto":{"number":3}},"method":{"description":"method of curl GET, POST, PUT, PATCH, DELETE","type":"string","__proto":{"number":4}},"href":{"description":"link","type":"string","__proto":{"number":5}},"type":{"description":"mime type","type":"string","__proto":{"number":6}}}},"google.type.date":{"name":"Date","type":"google.type.date","description":"Date, https://github.com/googleapis/googleapis/blob/master/google/type/date.proto ","fields":{"display_name":{"description":"Localized String representation of date","type":"string","meta":{"label":"Datum","default":"","hint":"","readonly":true},"constraints":{},"options":[],"__proto":{"number":4}},"year":{"description":"Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year.","type":"int","meta":{"default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":1,"type":"int32"}},"month":{"description":"Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day.","type":"int","meta":{"default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":2,"type":"int32"}},"day":{"description":"Day of month. Must be from 1 to 31 and valid for the year and month, or 0. if specifying a year by itself or a year and month where the day is not significant.","type":"int","meta":{"default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":3,"type":"int32"}}}},"google.type.money":{"name":"Money","type":"google.type.money","description":"Represents an amount of money with its currency type. https://github.com/googleapis/googleapis/blob/master/google/type/money.proto","fields":{"display_name":{"description":"String representation of money entity","type":"string","meta":{"default":"","hint":"","readonly":true},"constraints":{},"options":[],"__proto":{"number":1}},"currency_code":{"description":"The 3-letter currency code defined in ISO 4217.","type":"string","meta":{"label":"Währungscode","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":2}},"units":{"description":"The whole units of the amount.","type":"int","meta":{"label":"Ganzahliger Währungsbetrag","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":3,"type":"int64"}},"nanos":{"description":"Number of nano (10^-9) units of the amount. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.","type":"int","meta":{"label":"Nanos","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":4,"type":"int64"}}}},"person":{"name":"Person","type":"person","description":"Person message type","fields":{"display_name":{"description":"Localized String representation of a person","type":"string","meta":{"label":"Person","default":"","hint":"","readonly":true},"constraints":{},"options":[],"__proto":{"number":1}},"name":{"description":"Name of a person","type":"string","meta":{"label":"Name","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":2}},"first_name":{"description":"First name of a person","type":"string","meta":{"label":"First name","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":3}},"phone_nr":{"description":"Internal phone number","type":"string","meta":{"label":"Phone No","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":4}},"skills":{"description":"List of main skills of a person","type":"string","meta":{"label":"Skills","default":"","hint":"","repeated":true},"constraints":{},"options":[],"__proto":{"number":5}}}},"person_collection":{"name":"PersonCollection","type":"person_collection","mime":"application/person_collection+json","description":"Collection with persons inside","fields":{"meta":{"description":"Meta for the response","type":"furo.meta","__proto":{"number":2}},"links":{"description":"Hateoas links","type":"furo.link","meta":{"repeated":true},"__proto":{"number":3}},"entities":{"description":"person entities","type":"person","meta":{"repeated":true},"__proto":{"number":4}}}},"project":{"name":"Project","type":"project","description":"Project description","fields":{"display_name":{"description":"Localized String representation of a project","type":"string","meta":{"label":"Project","default":"","hint":"","readonly":true},"constraints":{},"options":[],"__proto":{"number":1}},"start":{"description":"Start date of the project","type":"google.type.date","meta":{"label":"Project start","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":2,"type":"google.type.date"}},"end":{"description":"Prospective end date of the project","type":"google.type.date","meta":{"label":"Project end","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":3,"type":"google.type.date"}},"description":{"description":"Short project description","type":"string","meta":{"label":"Description","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":4}},"members":{"description":"List of project members","type":"person","meta":{"label":"Project members","default":"","hint":"","repeated":true},"constraints":{},"options":[],"__proto":{"number":5,"type":"person"}},"cost_limit":{"description":"Project cost limit","type":"google.type.money","meta":{"label":"Cost limit","default":"","hint":"","required":true},"constraints":{"max":25000},"options":[],"__proto":{"number":6,"type":"google.type.money"}}}},"task":{"name":"Task","type":"task","description":"Task data description","fields":{"display_name":{"description":"Localized String representation of a task","type":"string","meta":{"label":"Task","default":"","hint":"","readonly":true},"constraints":{},"options":[],"__proto":{"number":1}},"description":{"description":"Short task description","type":"string","meta":{"label":"Description","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":2}},"estimated_time":{"description":"Estimated time in days","type":"int","meta":{"label":"Est. days","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":3,"type":"int32"}},"owner":{"description":"Owner of a task","type":"person","meta":{"label":"Owner","default":"","hint":""},"constraints":{},"options":[],"__proto":{"number":4,"type":"person"}},"subtasks":{"description":"List of subtasks","type":"task","meta":{"label":"Subtask","default":"","hint":"","required":true},"constraints":{},"options":[],"__proto":{"number":5,"type":"task"}}}},"tree":{"name":"Tree","type":"tree","description":"Singletonresource of the navigationtree","fields":{"display_name":{"description":"String representation of the node","type":"string","meta":{"readonly":true,"tree-searcb-index":true},"__proto":{"number":1}},"id":{"description":"Id of the node","type":"string","__proto":{"number":2}},"subtitle":{"description":"subtitle of the node","type":"string","meta":{"tree-searcb-index":true},"__proto":{"number":3}},"description":{"description":"description of the node","meta":{"tree-searcb-index":true},"type":"string","__proto":{"number":4}},"panel":{"description":"Indicator which panel type is loaded","type":"string","meta":{"readonly":true,"tree-searcb-index":false},"__proto":{"number":5}},"icon":{"description":"icon of the node","type":"string","__proto":{"number":6}},"key_words":{"description":"key words of the node","meta":{"tree-searcb-index":true},"type":"string","__proto":{"number":7}},"has_error":{"description":"if node has error","type":"bool","__proto":{"number":8}},"open":{"description":"node open or not","type":"bool","__proto":{"number":9}},"root":{"description":"Rootnode of the tree","type":"treeitem","meta":{},"__proto":{"number":10}}}},"treeitem":{"name":"Treeitem","type":"treeitem","description":"Item of the navigationtree","fields":{"display_name":{"description":"String representation of the node","type":"string","meta":{"readonly":true,"tree-searcb-index":true},"__proto":{"number":1}},"id":{"description":"Id of the node","type":"string","__proto":{"number":2}},"subtitle":{"description":"subtitle of the node","type":"string","meta":{"tree-searcb-index":true},"__proto":{"number":3}},"description":{"description":"description of the node","meta":{"tree-searcb-index":true},"type":"string","__proto":{"number":4}},"icon":{"description":"icon of the node","type":"string","__proto":{"number":5}},"key_words":{"description":"key words of the node","meta":{"tree-searcb-index":true},"type":"string","__proto":{"number":6}},"has_error":{"description":"if node has error","type":"bool","__proto":{"number":7}},"open":{"description":"node open or not","type":"bool","__proto":{"number":8}},"link":{"description":"Deeplink information of this node","type":"vnd.com.adcubum.link","__proto":{"number":9}},"children":{"description":"Children of this node","type":"vnd.com.adcubum.treeitem","meta":{"repeated":true},"__proto":{"number":10}}}}}
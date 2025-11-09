package core

const (
	DocxBlockTypePage              int = 1   // 文档 Block
	DocxBlockTypeText              int = 2   // 文本 Block
	DocxBlockTypeHeading1          int = 3   // 一级标题 Block
	DocxBlockTypeHeading2          int = 4   // 二级标题 Block
	DocxBlockTypeHeading3          int = 5   // 三级标题 Block
	DocxBlockTypeHeading4          int = 6   // 四级标题 Block
	DocxBlockTypeHeading5          int = 7   // 五级标题 Block
	DocxBlockTypeHeading6          int = 8   // 六级标题 Block
	DocxBlockTypeHeading7          int = 9   // 七级标题 Block
	DocxBlockTypeHeading8          int = 10  // 八级标题 Block
	DocxBlockTypeHeading9          int = 11  // 九级标题 Block
	DocxBlockTypeBullet            int = 12  // 无序列表 Block
	DocxBlockTypeOrdered           int = 13  // 有序列表 Block
	DocxBlockTypeCode              int = 14  // 代码块 Block
	DocxBlockTypeQuote             int = 15  // 引用 Block
	DocxBlockTypeTodo              int = 17  // 任务 Block
	DocxBlockTypeBitable           int = 18  // 多维表格 Block
	DocxBlockTypeCallout           int = 19  // 高亮块 Block
	DocxBlockTypeChatCard          int = 20  // 群聊卡片 Block
	DocxBlockTypeDiagram           int = 21  // 流程图/UML Block
	DocxBlockTypeDivider           int = 22  // 分割线 Block
	DocxBlockTypeFile              int = 23  // 文件 Block
	DocxBlockTypeGrid              int = 24  // 分栏 Block
	DocxBlockTypeGridColumn        int = 25  // 分栏列 Block
	DocxBlockTypeIframe            int = 26  // 内嵌 Block
	DocxBlockTypeImage             int = 27  // 图片 Block
	DocxBlockTypeISV               int = 28  // 三方 Block
	DocxBlockTypeMindnote          int = 29  // 思维笔记 Block
	DocxBlockTypeSheet             int = 30  // 电子表格 Block
	DocxBlockTypeTable             int = 31  // 表格 Block
	DocxBlockTypeTableCell         int = 32  // 单元格 Block
	DocxBlockTypeView              int = 33  // 视图 Block
	DocxBlockTypeQuoteContainer    int = 34  // 引用容器 Block
	DocxBlockTypeTask              int = 35  // 任务 Block
	DocxBlockTypeOkr               int = 36  // OKR Block
	DocxBlockTypeOkrObjective      int = 37  // OKR 目标 Block
	DocxBlockTypeOkrKeyResult      int = 38  // OKR 关键结果 Block
	DocxBlockTypeOkrProgress       int = 39  // OKR 进展 Block
	DocxBlockTypeAddOns            int = 40  // 新版文档小组件 Block
	DocxBlockTypeJiraIssue         int = 41  // Jira Issue Block
	DocxBlockTypeWikiCatalog       int = 42  // Wiki子页面 Block（旧版）
	DocxBlockTypeBoard             int = 43  // 画板 Block
	DocxBlockTypeAgenda            int = 44  // 议程 Block
	DocxBlockTypeAgendaItem        int = 45  // 议程项 Block
	DocxBlockTypeAgendaItemTitle   int = 46  // 议程项标题 Block
	DocxBlockTypeAgendaItemContent int = 47  // 议程项内容 Block
	DocxBlockTypeLinkPreview       int = 48  // 链接预览 Block
	DocxBlockTypeSourceSynced      int = 49  // 源同步 Block
	DocxBlockTypeReferenceSynced   int = 50  // 引用同步 Block
	DocxBlockTypeSubPageList       int = 51  // Wiki 子页面列表 Block（新版）
	DocxBlockTypeAiTemplate        int = 52  // AI 模板 Block
	DocxBlockTypeUndefined         int = 999 // 未支持 Block
)

const (
	DocxCodeLanguagePlainText             int = 1
	DocxCodeLanguageABAP                  int = 2
	DocxCodeLanguageAda                   int = 3
	DocxCodeLanguageApache                int = 4
	DocxCodeLanguageApex                  int = 5
	DocxCodeLanguageAssembly              int = 6
	DocxCodeLanguageBash                  int = 7
	DocxCodeLanguageCSharp                int = 8
	DocxCodeLanguageCPlusPlus             int = 9
	DocxCodeLanguageC                     int = 10
	DocxCodeLanguageCOBOL                 int = 11
	DocxCodeLanguageCSS                   int = 12
	DocxCodeLanguageCoffeeScript          int = 13
	DocxCodeLanguageD                     int = 14
	DocxCodeLanguageDart                  int = 15
	DocxCodeLanguageDelphi                int = 16
	DocxCodeLanguageDjango                int = 17
	DocxCodeLanguageDockerfile            int = 18
	DocxCodeLanguageErlang                int = 19
	DocxCodeLanguageFortran               int = 20
	DocxCodeLanguageFoxPro                int = 21
	DocxCodeLanguageGo                    int = 22
	DocxCodeLanguageGroovy                int = 23
	DocxCodeLanguageHTML                  int = 24
	DocxCodeLanguageHTMLBars              int = 25
	DocxCodeLanguageHTTP                  int = 26
	DocxCodeLanguageHaskell               int = 27
	DocxCodeLanguageJSON                  int = 28
	DocxCodeLanguageJava                  int = 29
	DocxCodeLanguageJavaScript            int = 30
	DocxCodeLanguageJulia                 int = 31
	DocxCodeLanguageKotlin                int = 32
	DocxCodeLanguageLateX                 int = 33
	DocxCodeLanguageLisp                  int = 34
	DocxCodeLanguageLogo                  int = 35
	DocxCodeLanguageLua                   int = 36
	DocxCodeLanguageMATLAB                int = 37
	DocxCodeLanguageMakefile              int = 38
	DocxCodeLanguageMarkdown              int = 39
	DocxCodeLanguageNginx                 int = 40
	DocxCodeLanguageObjective             int = 41
	DocxCodeLanguageOpenEdgeABL           int = 42
	DocxCodeLanguagePHP                   int = 43
	DocxCodeLanguagePerl                  int = 44
	DocxCodeLanguagePostScript            int = 45
	DocxCodeLanguagePowerShell            int = 46
	DocxCodeLanguageProlog                int = 47
	DocxCodeLanguageProtoBuf              int = 48
	DocxCodeLanguagePython                int = 49
	DocxCodeLanguageR                     int = 50
	DocxCodeLanguageRPG                   int = 51
	DocxCodeLanguageRuby                  int = 52
	DocxCodeLanguageRust                  int = 53
	DocxCodeLanguageSAS                   int = 54
	DocxCodeLanguageSCSS                  int = 55
	DocxCodeLanguageSQL                   int = 56
	DocxCodeLanguageScala                 int = 57
	DocxCodeLanguageScheme                int = 58
	DocxCodeLanguageScratch               int = 59
	DocxCodeLanguageShell                 int = 60
	DocxCodeLanguageSwift                 int = 61
	DocxCodeLanguageThrift                int = 62
	DocxCodeLanguageTypeScript            int = 63
	DocxCodeLanguageVBScript              int = 64
	DocxCodeLanguageVisual                int = 65
	DocxCodeLanguageXML                   int = 66
	DocxCodeLanguageYAML                  int = 67
	DocxCodeLanguageCMAKE                 int = 68
	DocxCodeLanguageDiff                  int = 69
	DocxCodeLanguageGherkin               int = 70
	DocxCodeLanguageGraphQL               int = 71
	DocxCodeLanguageOpenGLShadingLanguage int = 72
	DocxCodeLanguageProperties            int = 73
	DocxCodeLanguageSolidity              int = 74
	DocxCodeLanguageTOML                  int = 75
)

var DocxCodeLang2MdStr = map[int]string{
	DocxCodeLanguagePlainText:             "",
	DocxCodeLanguageABAP:                  "abap",
	DocxCodeLanguageAda:                   "ada",
	DocxCodeLanguageApache:                "apache",
	DocxCodeLanguageApex:                  "apex",
	DocxCodeLanguageAssembly:              "asm",
	DocxCodeLanguageBash:                  "bash",
	DocxCodeLanguageCSharp:                "csharp",
	DocxCodeLanguageCPlusPlus:             "cpp",
	DocxCodeLanguageC:                     "c",
	DocxCodeLanguageCOBOL:                 "cobol",
	DocxCodeLanguageCSS:                   "css",
	DocxCodeLanguageCoffeeScript:          "coffee",
	DocxCodeLanguageD:                     "d",
	DocxCodeLanguageDart:                  "dart",
	DocxCodeLanguageDelphi:                "delphi",
	DocxCodeLanguageDjango:                "django",
	DocxCodeLanguageDockerfile:            "docker",
	DocxCodeLanguageErlang:                "erlang",
	DocxCodeLanguageFortran:               "fortran",
	DocxCodeLanguageFoxPro:                "foxpro",
	DocxCodeLanguageGo:                    "go",
	DocxCodeLanguageGroovy:                "groovy",
	DocxCodeLanguageHTML:                  "html",
	DocxCodeLanguageHTMLBars:              "htmlbars",
	DocxCodeLanguageHTTP:                  "http",
	DocxCodeLanguageHaskell:               "haskell",
	DocxCodeLanguageJSON:                  "json",
	DocxCodeLanguageJava:                  "java",
	DocxCodeLanguageJavaScript:            "javascript",
	DocxCodeLanguageJulia:                 "julia",
	DocxCodeLanguageKotlin:                "kotlin",
	DocxCodeLanguageLateX:                 "latex",
	DocxCodeLanguageLisp:                  "lisp",
	DocxCodeLanguageLogo:                  "logo",
	DocxCodeLanguageLua:                   "lua",
	DocxCodeLanguageMATLAB:                "matlab",
	DocxCodeLanguageMakefile:              "makefile",
	DocxCodeLanguageMarkdown:              "markdown",
	DocxCodeLanguageNginx:                 "nginx",
	DocxCodeLanguageObjective:             "objectivec",
	DocxCodeLanguageOpenEdgeABL:           "openedgeabl",
	DocxCodeLanguagePHP:                   "php",
	DocxCodeLanguagePerl:                  "perl",
	DocxCodeLanguagePostScript:            "postscript",
	DocxCodeLanguagePowerShell:            "powershell",
	DocxCodeLanguageProlog:                "prolog",
	DocxCodeLanguageProtoBuf:              "protobuf",
	DocxCodeLanguagePython:                "python",
	DocxCodeLanguageR:                     "r",
	DocxCodeLanguageRPG:                   "rpgle",
	DocxCodeLanguageRuby:                  "ruby",
	DocxCodeLanguageRust:                  "rust",
	DocxCodeLanguageSAS:                   "sas",
	DocxCodeLanguageSCSS:                  "scss",
	DocxCodeLanguageSQL:                   "sql",
	DocxCodeLanguageScala:                 "scala",
	DocxCodeLanguageScheme:                "scheme",
	DocxCodeLanguageScratch:               "scratch",
	DocxCodeLanguageShell:                 "shell",
	DocxCodeLanguageSwift:                 "swift",
	DocxCodeLanguageThrift:                "thrift",
	DocxCodeLanguageTypeScript:            "typescript",
	DocxCodeLanguageVBScript:              "vbscript",
	DocxCodeLanguageVisual:                "visual",
	DocxCodeLanguageXML:                   "xml",
	DocxCodeLanguageYAML:                  "yaml",
	DocxCodeLanguageCMAKE:                 "cmake",
	DocxCodeLanguageDiff:                  "diff",
	DocxCodeLanguageGherkin:               "gherkin",
	DocxCodeLanguageGraphQL:               "graphql",
	DocxCodeLanguageOpenGLShadingLanguage: "glsl",
	DocxCodeLanguageProperties:            "properties",
	DocxCodeLanguageSolidity:              "solidity",
	DocxCodeLanguageTOML:                  "toml",
}

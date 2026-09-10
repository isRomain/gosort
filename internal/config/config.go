package config

// Extensions par catégorie
var (
	extensionsOfApplications = []string{"app", "dmg", "exe", "msi", "deb", "rpm"}
	extensionsOfArchives     = []string{"zip", "rar", "7z", "tar", "gz", "tgz", "bz2", "xz", "iso"}
	extensionsOfVideos       = []string{"mp4", "mkv", "mov", "avi", "wmv", "flv", "webm", "mpeg", "mpg", "3gp"}
	extensionsOfAudios       = []string{"mp3", "wav", "flac", "aac", "ogg", "wma", "m4a", "opus", "aiff", "alac"}
	extensionsOfDocuments    = []string{"pdf", "doc", "docx", "odt", "rtf", "txt", "xls", "xlsx", "ods", "ppt", "pptx", "md"}
	extensionsOfImages       = []string{"png", "jpg", "jpeg", "gif", "bmp", "webp", "tiff", "svg", "heic", "ico", "avif", "raw", "psd", "icns"}

	extensionsOfScripts = []string{
		"py", "ts", "js", "sh", "rb", "pl", "php", "ps1", "sql", "rs", "bat", "c", "cpp",
		"h", "hpp", "java", "kt", "go", "cs", "swift", "dart", "lua", "r", "html", "css", "scss",
		"json", "yaml", "yml", "xml", "toml", "ini", "md", "jsx", "tsx", "mjs", "cjs", "zsh", "fish",
		"cmd", "vb", "fs", "scala", "groovy", "perl", "jl", "hs", "clj", "elm", "tex", "asm", "s", "vue",
	}
)

// Dossiers de destination
var (
	pathOfApplications = "/Users/romain/Downloads/Applications"
	pathOfArchives     = "/Users/romain/Downloads/Archives"
	pathOfVideos       = "/Users/romain/Movies/Videos"
	pathOfAudios       = "/Users/romain/Downloads/Audios"
	pathOfDocuments    = "/Users/romain/Downloads/Docs"
	pathOfImages       = "/Users/romain/Pictures/Images"
	pathOfScripts      = "/Users/romain/Downloads/Scripts"
)

// Association dossier -> extensions gérées
var MyDict = map[string][]string{
	pathOfApplications: extensionsOfApplications,
	pathOfArchives:     extensionsOfArchives,
	pathOfVideos:       extensionsOfVideos,
	pathOfAudios:       extensionsOfAudios,
	pathOfDocuments:    extensionsOfDocuments,
	pathOfImages:       extensionsOfImages,
	pathOfScripts:      extensionsOfScripts,
}

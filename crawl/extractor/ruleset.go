package extractor

const (
	NSPath uint = iota
	NSDataset
	NSCombine
)

type RuleSet struct {
	Collection string
	NameSpace  uint
	SRSText    string
	Proj4Text  string
	Pattern    string
}

const (
	SRSDetect string = ""
	SRSWGS84  string = `GEOGCS["WGS 84",DATUM["WGS_1984",SPHEROID["WGS 84",6378137,298.257223563,AUTHORITY["EPSG","7030"]],TOWGS84[0,0,0,0,0,0,0],AUTHORITY["EPSG","6326"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9108"]],AUTHORITY["EPSG","4326"]]`
)

const (
	Proj4Detect string = ""
	// Generated by proj4; unsure if the trailing space is significant
	Proj4WGS84 string = `+proj=longlat +ellps=WGS84 +towgs84=0,0,0,0,0,0,0 +no_defs `
)

var CollectionRuleSets = []RuleSet{
	RuleSet{
		"landsat",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`LC(?P<mission>\d)(?P<path>\d\d\d)(?P<row>\d\d\d)(?P<year>\d\d\d\d)(?P<julian_day>\d\d\d)(?P<processing_level>[a-zA-Z0-9]+)_(?P<band>[a-zA-Z0-9]+)`,
	},
	RuleSet{
		"modis43A4",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^LHTC_(?P<year>\d\d\d\d)(?P<julian_day>\d\d\d).(?P<horizontal>h\d\d)(?P<vertical>v\d\d).(?P<resolution>\d\d\d).[0-9]+`,
	},
	RuleSet{
		"lhtc",
		NSCombine,
		SRSDetect,
		Proj4Detect,
		`^COMPOSITE_(?P<namespace>LOW|HIGH).+_PER_20.nc$`,
	},
	RuleSet{
		"modis1",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^(?P<product>MCD\d\d[A-Z]\d).A(?P<year>\d\d\d\d)(?P<julian_day>\d\d\d).(?P<horizontal>h\d\d)(?P<vertical>v\d\d).(?P<resolution>\d\d\d).[0-9]+`,
	},
	RuleSet{
		"modis-fc",
		NSPath,
		SRSDetect,
		Proj4Detect,
		`^(?P<product>FC).v302.(?P<collection>MCD43A4).h(?P<horizontal>\d\d)v(?P<vertical>\d\d).(?P<year>\d\d\d\d).(?P<resolution>\d\d\d).(?P<namespace>[A-Z0-9]+).jp2$`,
	},
	RuleSet{
		"modis2",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`M(?P<satellite>[OD|YD])(?P<product>[0-9]+_[A-Z0-9]+).A[0-9]+.[0-9]+.(?P<collection_version>\d\d\d).(?P<year>\d\d\d\d)(?P<julian_day>\d\d\d)(?P<hour>\d\d)(?P<minute>\d\d)(?P<second>\d\d)`,
	},
	RuleSet{
		"modisJP",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^(?P<product>FC).v302.(?P<root_product>MCD\d\d[A-Z]\d).h(?P<horizontal>\d\d)v(?P<vertical>\d\d).(?P<year>\d\d\d\d).(?P<resolution>\d\d\d).`,
	},
	RuleSet{
		"sentinel2",
		NSPath,
		SRSDetect,
		Proj4Detect,
		`^T(?P<zone>\d\d)(?P<sensor>[A-Z]+)_(?P<year>\d\d\d\d)(?P<month>\d\d)(?P<day>\d\d)T(?P<hour>\d\d)(?P<minute>\d\d)(?P<second>\d\d)_(?P<namespace>B\d\d).jp2$`,
	},
	RuleSet{
		"modisJP_LR",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^(?P<product>FC_LR).v302.(?P<root_product>MCD\d\d[A-Z]\d).h(?P<horizontal>\d\d)v(?P<vertical>\d\d).(?P<year>\d\d\d\d).(?P<resolution>\d\d\d).`,
	},
	RuleSet{
		"himawari8",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^(?P<year>\d\d\d\d)(?P<month>\d\d)(?P<day>\d\d)(?P<hour>\d\d)(?P<minute>\d\d)(?P<second>\d\d)-P1S-(?P<product>ABOM[0-9A-Z_]+)-PRJ_GEOS141_(?P<resolution>\d+)-HIMAWARI8-AHI`,
	},
	RuleSet{
		"agdc_landsat1",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`LS(?P<mission>\d)_(?P<sensor>[A-Z]+)_(?P<correction>[A-Z]+)_(?P<epsg>\d+)_(?P<x_coord>-?\d+)_(?P<y_coord>-?\d+)_(?P<year>\d\d\d\d).`,
	},
	RuleSet{
		"elevation_ga",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`^Elevation_1secSRTM_DEMs_v1.0_DEM-S_Tiles_e(?P<longitude>\d+)s(?P<latitude>\d+)dems.nc$`,
	},
	RuleSet{
		"chirps2.0",
		NSDataset,
		SRSWGS84,
		Proj4WGS84,
		`^chirps-v2.0.(?P<year>\d\d\d\d).dekads.nc$`,
	},
	RuleSet{
		"era-interim",
		NSPath,
		SRSDetect,
		Proj4Detect,
		`^(?P<namespace>[a-z0-9]+)_(?P<accum>\dhrs)_ERAI_historical_(?P<levels>[a-z\-]+)_(?P<start_year>\d\d\d\d)(?P<start_month>\d\d)(?P<start_day>\d\d)_(?P<end_year>\d\d\d\d)(?P<end_month>\d\d)(?P<end_day>\d\d).nc$`,
	},
	RuleSet{
		"agdc_landsat2",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`LS(?P<mission>\d)_OLI_(?P<sensor>[A-Z]+)_(?P<product>[A-Z]+)_(?P<epsg>\d+)_(?P<x_coord>-?\d+)_(?P<y_coord>-?\d+)_(?P<year>\d\d\d\d).`,
	},
	RuleSet{
		"agdc_dem",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`SRTM_(?P<product>[A-Z]+)_(?P<x_coord>-?\d+)_(?P<y_coord>-?\d+)_(?P<year>\d\d\d\d)(?P<month>\d\d)(?P<day>\d\d)(?P<hour>\d\d)(?P<minute>\d\d)(?P<second>\d\d)`,
	},
	RuleSet{
		"default",
		NSDataset,
		SRSDetect,
		Proj4Detect,
		`.+`,
	},
}

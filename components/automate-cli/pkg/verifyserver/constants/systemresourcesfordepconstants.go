package constants

const (
	MIN_CPU_COUNT = 2
	MIN_CPU_SPEED = 2.0
	MIN_MEMORY    = 7

	HAB_FREE_DISK_BEFORE_DEP_A2      = 80
	HAB_FREE_DISK_BEFORE_DEP_CS      = 80
	HAB_FREE_DISK_BEFORE_DEP_PG      = 150
	HAB_FREE_DISK_BEFORE_DEP_OS      = 150
	HAB_FREE_DISK_BEFORE_DEP_BASTION = 150

	HAB_FREE_DISK_AFTER_DEP_A2      = 10
	HAB_FREE_DISK_AFTER_DEP_CS      = 10
	HAB_FREE_DISK_AFTER_DEP_PG      = 50
	HAB_FREE_DISK_AFTER_DEP_OS      = 50
	HAB_FREE_DISK_AFTER_DEP_BASTION = 30

	HAB_FREE_DISK_AFTER_DEP_A2_IN_PER      = 0.05
	HAB_FREE_DISK_AFTER_DEP_CS_IN_PER      = 0.05
	HAB_FREE_DISK_AFTER_DEP_PG_IN_PER      = 0.15
	HAB_FREE_DISK_AFTER_DEP_OS_IN_PER      = 0.15
	HAB_FREE_DISK_AFTER_DEP_BASTION_IN_PER = 0.15

	TMP_FREE_DISK_IN_GB  = 10
	TMP_FREE_DISK_IN_PER = 0.05

	ROOT_FREE_DISK_IN_GB  = 20
	ROOT_FREE_DISK_IN_PER = 0.20

	GB_TO_BYTES = 1024 * 1024 * 1024

	TMP_DIR_REQUIRED_PERMISSION = "-rwxrwxrwx"
)

const (
	CPU_COUNT_CHECK_TITLE   = "CPU count check"
	CPU_SPEED_CHECK_TITLE   = "CPU speed check"
	MEMORY_SIZE_CHECK_TITLE = "Memory size check"
	FREE_SPACE_CHECK        = "%s free space check"

	SUCCESS_MSG        = "%s should have free space >=%vGB"
	ERROR_MSG          = "%s free space is %0.2fGB"
	SUCCESS_MSG_IN_PER = " or %v%% of total size of /hab"

	RESOLUTION_MSG = "Please run system on supported platform"

	PERMISSION_CHECK       = "%s permission check"
	PERMISSION_SUCCESS_MSG = "%s should have %s permission"
	PERMISSION_ERROR_MSG   = "%s permission is %s"
)

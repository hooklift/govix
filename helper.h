#ifndef helpers_h
#define helpers_h 1

VixError getHandle(
	VixHandle jobHandle,
	VixPropertyID prop1,
	VixHandle* handle,
	VixPropertyID prop2);

VixError allocVmPasswordPropList(
	VixHandle handle,
	VixHandle* resultHandle,
	char* password);

VixError getScreenshotBytes(
	VixHandle handle,
	int* byte_count,
	char* screen_bits);

VixError runProgramResult(
	VixHandle jobHandle,
	uint64* pid,
	int* elapsedTime,
	int* exitCode);

VixError getSharedFolder(
	VixHandle jobHandle,
	char* folderName,
	char* folderHostPath,
	int* folderFlags);

VixError getFileInfo(VixHandle jobHandle,
					 int64* fsize,
					 int* flags,
					 int64* modtime);

void find_items_callback(
	VixHandle jobHandle,
	VixEventType eventType,
	VixHandle moreEventInfo,
	void *items);

VixError getNumSharedFolders(VixHandle jobHandle, int* numSharedFolders);
VixError readVariable(VixHandle jobHandle, char* readValue);
VixError getTempFilePath(VixHandle jobHandle, char* tempFilePath);
VixError isFileOrDir(VixHandle jobHandle, int* result);
VixError getVMUrl(char* url, VixHandle moreEvtInfo);

#endif

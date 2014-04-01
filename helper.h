#ifndef helpers_h
#define helpers_h 1

VixError get_vix_handle(
	VixHandle jobHandle,
	VixPropertyID prop1,
	VixHandle* handle,
	VixPropertyID prop2);

VixError alloc_vm_pwd_proplist(
	VixHandle handle,
	VixHandle* resultHandle,
	char* password);

VixError get_screenshot_bytes(
	VixHandle handle,
	int* byte_count,
	char* screen_bits);

VixError get_program_output(
	VixHandle jobHandle,
	uint64* pid,
	int* elapsedTime,
	int* exitCode);

VixError get_shared_folder(
	VixHandle jobHandle,
	char* folderName,
	char* folderHostPath,
	int* folderFlags);

VixError get_file_info(VixHandle jobHandle,
					 int64* fsize,
					 int* flags,
					 int64* modtime);

void find_items_callback(
	VixHandle jobHandle,
	VixEventType eventType,
	VixHandle moreEventInfo,
	void *items);

VixError get_num_shared_folders(VixHandle jobHandle, int* numSharedFolders);
VixError read_variable(VixHandle jobHandle, char* readValue);
VixError get_temp_filepath(VixHandle jobHandle, char* tempFilePath);
VixError is_file_or_dir(VixHandle jobHandle, int* result);
VixError get_vm_url(char* url, VixHandle moreEvtInfo);

#endif

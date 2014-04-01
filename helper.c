#include <stdio.h>
#include <string.h>
#include "vix.h"

VixError getHandle(
	VixHandle jobHandle,
	VixPropertyID prop1,
	VixHandle* handle,
	VixPropertyID prop2) {

	return VixJob_Wait(jobHandle,
		prop1,
		handle,
		prop2);
}

VixError allocVmPasswordPropList(
	VixHandle handle,
	VixHandle* resultHandle,
	char* password
) {
	return VixPropertyList_AllocPropertyList(handle,
                                        resultHandle,
                                        VIX_PROPERTY_VM_ENCRYPTION_PASSWORD,
                                        password,
                                        VIX_PROPERTY_NONE);
}

VixError getScreenshotBytes(
	VixHandle handle,
	int* byte_count,
	char* screen_bits) {

	return VixJob_Wait(handle,
		VIX_PROPERTY_JOB_RESULT_SCREEN_IMAGE_DATA,
		byte_count, screen_bits,
		VIX_PROPERTY_NONE);
}

VixError getNumSharedFolders(VixHandle jobHandle, int* numSharedFolders) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_COUNT,
		numSharedFolders,
		VIX_PROPERTY_NONE);
}

VixError readVariable(VixHandle jobHandle, char* readValue) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_VM_VARIABLE_STRING,
		readValue,
		VIX_PROPERTY_NONE);
}

VixError getTempFilePath(VixHandle jobHandle, char* tempFilePath) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_ITEM_NAME,
		tempFilePath,
		VIX_PROPERTY_NONE);
}

VixError isFileOrDir(VixHandle jobHandle, int* result) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_GUEST_OBJECT_EXISTS,
		result,
		VIX_PROPERTY_NONE);
}

VixError runProgramResult(
	VixHandle jobHandle,
	uint64* pid,
	int* elapsedTime,
	int* exitCode) {
	return VixJob_Wait(jobHandle,
		VIX_PROPERTY_JOB_RESULT_PROCESS_ID,
		pid,
		VIX_PROPERTY_JOB_RESULT_GUEST_PROGRAM_ELAPSED_TIME,
		elapsedTime,
		VIX_PROPERTY_JOB_RESULT_GUEST_PROGRAM_EXIT_CODE,
		exitCode,
		VIX_PROPERTY_NONE);
}

VixError getSharedFolder(
	VixHandle jobHandle,
	char* folderName,
	char* folderHostPath,
	int* folderFlags) {

	return VixJob_Wait(jobHandle,
    	VIX_PROPERTY_JOB_RESULT_ITEM_NAME, folderName,
        VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_HOST, folderHostPath,
        VIX_PROPERTY_JOB_RESULT_SHARED_FOLDER_FLAGS, folderFlags,
		VIX_PROPERTY_NONE);
}


VixError getVMUrl(char* url, VixHandle moreEvtInfo) {
	return	Vix_GetProperties(	moreEvtInfo,
								VIX_PROPERTY_FOUND_ITEM_LOCATION,
								url,
								VIX_PROPERTY_NONE);
}

VixError getFileInfo(VixHandle jobHandle,
					 int64* fsize,
					 int* flags,
					 int64* modtime) {

	return Vix_GetProperties(jobHandle,
		                VIX_PROPERTY_JOB_RESULT_FILE_SIZE,
                        fsize,
                        VIX_PROPERTY_JOB_RESULT_FILE_FLAGS,
                        flags,
                        VIX_PROPERTY_JOB_RESULT_FILE_MOD_TIME,
                        modtime,
                        VIX_PROPERTY_NONE);
}

void find_items_callback(
	VixHandle jobHandle,
	VixEventType eventType,
	VixHandle moreEventInfo,
	void *items)
{
   VixError err = VIX_OK;
   char* url = NULL;

   	//Check callback event; ignore progress reports.
   	if (VIX_EVENTTYPE_FIND_ITEM != eventType) {
   		return;
   	}

  	// Found a virtual machine.
  	err = Vix_GetProperties(moreEventInfo,
                           VIX_PROPERTY_FOUND_ITEM_LOCATION,
                           &url,
                           VIX_PROPERTY_NONE);
   	if (VIX_OK != err) {
   		// Handle the error...
	  	printf("Error %s\n", Vix_GetErrorText(err, NULL));
   	}

   	if (items == NULL) {
   		printf("Allocating space for items...");
		items = (char**) malloc(sizeof(char) * 10);
   	}

   	strncpy(items, url, strlen(url));

	Vix_FreeBuffer(url);
	//printf("\nFound virtual machine: %s\n", items);
   	items++;
}

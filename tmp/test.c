#include "../vix.h"
#include <stdio.h>

// cc -L /Users/camilo/Dropbox/Development/go/src/github.com/c4milo/govix -lvixAllProducts -ldl -lpthread tmp/test.c -o test
static VixHandle hostHandle = VIX_INVALID_HANDLE;

void VixDiscoveryProc(VixHandle jobHandle,
                      VixEventType eventType,
                      VixHandle moreEventInfo,
                      void *clientData)
{
   VixError err = VIX_OK;
   char *url = NULL;

   // Check callback event; ignore progress reports.
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
      goto abort;
   }

   printf("\nFound virtual machine: %s", url);

abort:
   Vix_FreeBuffer(url);
}

int main()
{
   VixHandle jobHandle = VIX_INVALID_HANDLE;
   VixHandle hostHandle = VIX_INVALID_HANDLE;
   VixError err;
   jobHandle = VixHost_Connect(VIX_API_VERSION,
                               VIX_SERVICEPROVIDER_VMWARE_WORKSTATION,
                               NULL, // hostName
                               0, // hostPort
                               NULL, // userName
                               NULL, // password,
                               0, // options
                               VIX_INVALID_HANDLE, // propertyListHandle
                               NULL, // callbackProc
                               NULL); // clientData
   err = VixJob_Wait(jobHandle,
                     VIX_PROPERTY_JOB_RESULT_HANDLE,
                     &hostHandle,
                     VIX_PROPERTY_NONE);
   if (VIX_FAILED(err)) {
      // Handle the error...
      goto abort;
   }

   Vix_ReleaseHandle(jobHandle);

   printf("\nLooking for running virtual machines...\n");

   jobHandle = VixHost_FindItems(hostHandle,
                                 VIX_FIND_RUNNING_VMS,
                                 VIX_INVALID_HANDLE, // searchCriteria
                                 -1, // timeout
                                 VixDiscoveryProc,
                                 NULL);
   err = VixJob_Wait(jobHandle, VIX_PROPERTY_NONE);
   if (VIX_FAILED(err)) {
      // Handle the error...
      goto abort;
   }
abort:
   Vix_ReleaseHandle(jobHandle);
   VixHost_Disconnect(hostHandle);
}
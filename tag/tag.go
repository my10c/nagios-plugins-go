// Copyright (c) 2014 - 2017 badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//   * Redistributions of source code must retain the above copyright
//   notice, this list of conditions and the following disclaimer.
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//   * Neither the name of the <organization> nor the
//   names of its contributors may be used to endorse or promote products
//   derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Version		:	0.1
//
// Date			:	June 2, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 2, 2017	LIS			First release
//

package tag

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	myGlobal	"github.com/my10c/nagios-plugins-go/global"
)

// Function to get a tag specify in the file with the specify tagkey
func GetTagInfo() (string, error) {
	var tagInfo string
	var err error = nil

	// make sure both tagfile and tagkeyname were set
	if len(myGlobal.DefaultValues["tagfile"]) > 0 &&
		len(myGlobal.DefaultValues["tagkeyname"]) > 0 {
		// open given tag file
		tagFile, ok := os.Open(myGlobal.DefaultValues["tagfile"])
		if ok != nil {
			err = fmt.Errorf("Unable to open the tag file %s", myGlobal.DefaultValues["tagfile"])
			return tagInfo, err
		}
		// make sure we closed the file
		defer tagFile.Close()
		// now read file and search for the tagkeyname
		scanner := bufio.NewScanner(tagFile)
		for scanner.Scan() {
			currLine := scanner.Text()
			if strings.HasPrefix(currLine, myGlobal.DefaultValues["tagkeyname"]) {
				tagInfo = strings.TrimPrefix(currLine, myGlobal.DefaultValues["tagkeyname"])
				return strings.TrimSpace(tagInfo), err
			}
		}
	} else {
		err = fmt.Errorf("Missing either tagfile or tagkeyname or both")
		return tagInfo, err
	}
	err = fmt.Errorf("Requested tagkeyname %s not found", myGlobal.DefaultValues["tagkeyname"])
	return tagInfo, err
}

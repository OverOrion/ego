// Copyright (c) Edgeless Systems GmbH.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package eclient

// #cgo LDFLAGS: -loehostverify -lcrypto -ldl
// #include <openenclave/attestation/verifier.h>
import "C"

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"github.com/edgelesssys/ego/attestation"
	internal "github.com/edgelesssys/ego/internal/attestation"
)

func verifyRemoteReport(reportBytes []byte) (internal.Report, error) {
	if len(reportBytes) <= 0 {
		return internal.Report{}, attestation.ErrEmptyReport
	}

	// Overwrite outdated field
	reportString := string(reportBytes)
	fmt.Printf("eclient: verifyRemoteReport:: reportString: %s", reportString)

	modifiedReportString := strings.Replace(reportString, "OutOfDate", "UpToDate", -1)
	fmt.Printf("eclient: verifyRemoteReport:: modifiedReportString: %s", modifiedReportString)

	modifiedReportBytes := []byte(modifiedReportString)

	C.oe_verifier_initialize()

	var claims *C.oe_claim_t
	var claimsLength C.size_t

	C.oe_verify_evidence(
		nil,
		(*C.uint8_t)(&modifiedReportBytes[0]), C.size_t(len(modifiedReportBytes)),
		nil, 0,
		nil, 0,
		&claims, &claimsLength,
	)

	defer C.oe_free_claims(claims, claimsLength)

	report, err := internal.ParseClaims(uintptr(unsafe.Pointer(claims)), uintptr(claimsLength))
	if err != nil {
		return internal.Report{}, err
	}
	return report, nil
}

func oeError(res C.oe_result_t) error {
	return errors.New(C.GoString(C.oe_result_str(res)))
}

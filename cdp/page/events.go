package page

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/runtime"
)

// EventDomContentEventFired [no description].
type EventDomContentEventFired struct {
	Timestamp float64 `json:"timestamp,omitempty"`
}

// EventLoadEventFired [no description].
type EventLoadEventFired struct {
	Timestamp float64 `json:"timestamp,omitempty"`
}

// EventFrameAttached fired when frame has been attached to its parent.
type EventFrameAttached struct {
	FrameID       cdp.FrameID         `json:"frameId,omitempty"`       // Id of the frame that has been attached.
	ParentFrameID cdp.FrameID         `json:"parentFrameId,omitempty"` // Parent frame identifier.
	Stack         *runtime.StackTrace `json:"stack,omitempty"`         // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
}

// EventFrameNavigated fired once navigation of the frame has completed.
// Frame is now associated with the new loader.
type EventFrameNavigated struct {
	Frame *cdp.Frame `json:"frame,omitempty"` // Frame object.
}

// EventFrameDetached fired when frame has been detached from its parent.
type EventFrameDetached struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Id of the frame that has been detached.
}

// EventFrameStartedLoading fired when frame has started loading.
type EventFrameStartedLoading struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Id of the frame that has started loading.
}

// EventFrameStoppedLoading fired when frame has stopped loading.
type EventFrameStoppedLoading struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Id of the frame that has stopped loading.
}

// EventFrameScheduledNavigation fired when frame schedules a potential
// navigation.
type EventFrameScheduledNavigation struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Id of the frame that has scheduled a navigation.
	Delay   float64     `json:"delay,omitempty"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
}

// EventFrameClearedScheduledNavigation fired when frame no longer has a
// scheduled navigation.
type EventFrameClearedScheduledNavigation struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Id of the frame that has cleared its scheduled navigation.
}

// EventFrameResized [no description].
type EventFrameResized struct{}

// EventJavascriptDialogOpening fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) is about to open.
type EventJavascriptDialogOpening struct {
	Message string     `json:"message,omitempty"` // Message that will be displayed by the dialog.
	Type    DialogType `json:"type,omitempty"`    // Dialog type.
}

// EventJavascriptDialogClosed fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) has been closed.
type EventJavascriptDialogClosed struct {
	Result bool `json:"result,omitempty"` // Whether dialog was confirmed.
}

// EventScreencastFrame compressed image data requested by the
// startScreencast.
type EventScreencastFrame struct {
	Data      string                   `json:"data,omitempty"`      // Base64-encoded compressed image.
	Metadata  *ScreencastFrameMetadata `json:"metadata,omitempty"`  // Screencast frame metadata.
	SessionID int64                    `json:"sessionId,omitempty"` // Frame number.
}

// EventScreencastVisibilityChanged fired when the page with currently
// enabled screencast was shown or hidden .
type EventScreencastVisibilityChanged struct {
	Visible bool `json:"visible,omitempty"` // True if the page is visible.
}

// EventInterstitialShown fired when interstitial page was shown.
type EventInterstitialShown struct{}

// EventInterstitialHidden fired when interstitial page was hidden.
type EventInterstitialHidden struct{}

// EventNavigationRequested fired when a navigation is started if navigation
// throttles are enabled. The navigation will be deferred until
// processNavigation is called.
type EventNavigationRequested struct {
	IsInMainFrame bool   `json:"isInMainFrame,omitempty"` // Whether the navigation is taking place in the main frame or in a subframe.
	IsRedirect    bool   `json:"isRedirect,omitempty"`    // Whether the navigation has encountered a server redirect or not.
	NavigationID  int64  `json:"navigationId,omitempty"`
	URL           string `json:"url,omitempty"` // URL of requested navigation.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventPageDomContentEventFired,
	cdp.EventPageLoadEventFired,
	cdp.EventPageFrameAttached,
	cdp.EventPageFrameNavigated,
	cdp.EventPageFrameDetached,
	cdp.EventPageFrameStartedLoading,
	cdp.EventPageFrameStoppedLoading,
	cdp.EventPageFrameScheduledNavigation,
	cdp.EventPageFrameClearedScheduledNavigation,
	cdp.EventPageFrameResized,
	cdp.EventPageJavascriptDialogOpening,
	cdp.EventPageJavascriptDialogClosed,
	cdp.EventPageScreencastFrame,
	cdp.EventPageScreencastVisibilityChanged,
	cdp.EventPageInterstitialShown,
	cdp.EventPageInterstitialHidden,
	cdp.EventPageNavigationRequested,
}

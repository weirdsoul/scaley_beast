; range produces a range of integer numbers with the specified stepping.
(define (range first last step)
  (if (>= first last)
      '()
      (cons first (range (+ first step) last step))
  )
)

; circ calculates the coordinates on a circle with the specified
; center and radius at deg degrees. Degrees are defined clockwise,
; starting at the topmost point of the circle.
(define (circ centerX centerY radius deg)
  (let* (
       (rad (* (/ (- 90 deg) 360) 2 *pi*))
       (y (- centerY (* (sin rad) radius)))
       (x (+ centerX (* (cos rad) radius)))
       )
    (list x y)
  )
)

; Draw a single scale marking at the position specified by a circle around
; the specified center at deg degrees with the specified length.
(define (drawScale inLayer centerX centerY deg len padding)
  (let* (
	 (pencil_line (cons-array 4 'double))
         (ci (circ centerX centerY (- centerX (+ padding len)) deg))
         (co (circ centerX centerY (- centerX padding) deg))
	 )
    (aset pencil_line 0 (car ci))
    (aset pencil_line 1 (cadr ci))
    (aset pencil_line 2 (car co))
    (aset pencil_line 3 (cadr co))
    (gimp-pencil inLayer 4 pencil_line)    
  )
)

; Maps the specified input to a polynomial. In the typical case, this is a non-linear
; function to stretch part of the scale.
(define (map-to-polynomial x)
  (+ (* 1.194e-6 (expt x 4)) (* -0.00049614 (expt x 3))
     (* 0.062352 (expt x 2)) (* -0.21389 x) -0.86754)
)

(define (script-fu-gauge inImage inLayer step evenMarkerLen oddMarkerLen padding)
  (let*
      (
       (xRes (car(gimp-drawable-height inLayer)))
       (yRes (car(gimp-drawable-width inLayer)))
       (centerX (/ xRes 2))
       (centerY (/ yRes 2))
       (r (range 0 190 step))
       )
    (gimp-undo-push-group-start inImage)
    (while (not (null? r))
	   (drawScale inLayer centerX centerY (map-to-polynomial (car r))
		      (if (> (modulo (car r) (* step 2)) 0) oddMarkerLen evenMarkerLen)
		      padding)
	   (set! r (cdr r))
    )
    (gimp-undo-push-group-end inImage)
    (gimp-displays-flush) ; Now make sure to update the UI
  )
)
(script-fu-register
    "script-fu-gauge"                           ;func name
    "Draw gauge..."                             ;menu label
    "Creates a simple gauge."                   ;description
    "Andreas Eckleder"                          ;author
    "copyright 2017, Andreas Eckleder"          ;copyright notice
    "April 30, 2017"                            ;date created
    "RGBA"                                      ;image type
    SF-IMAGE       "Image"         0
    SF-DRAWABLE    "Drawable"      0
    SF-VALUE       "Step size"           "10"
    SF-VALUE       "Even marker length"  "10"
    SF-VALUE       "Odd marker length"   "5"
    SF-VALUE       "Padding"             "50"
)
(script-fu-menu-register "script-fu-gauge" "<Image>/Filters/Render")

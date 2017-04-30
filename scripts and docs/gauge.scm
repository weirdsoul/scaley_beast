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
       (rad (* (/ (modulo (- 90 deg) 360) 360) 2 *pi*))
       (y (- centerY (* (sin rad) radius)))
       (x (+ centerX (* (cos rad) radius)))
    )
    (list x y)
  )
)

; Draw a single scale marking at the position specified by a circle around
; the specified center at deg degrees with the specified length.
(define (drawScale inLayer centerX centerY deg len)
  (let* (
	 (pencil_line (cons-array 4 'double))
         (ci (circ centerX centerY (- centerX (+ 10 len)) deg))
         (co (circ centerX centerY (- centerX 10) deg))
	 )
    (aset pencil_line 0 (car ci))
    (aset pencil_line 1 (cadr ci))
    (aset pencil_line 2 (car co))
    (aset pencil_line 3 (cadr co))
    (gimp-pencil inLayer 4 pencil_line)    
  )
)

(define (script-fu-gauge inImage inLayer)
  (let*
      (
       (xRes (car(gimp-drawable-height inLayer)))
       (yRes (car(gimp-drawable-width inLayer)))
       (centerX (/ xRes 2))
       (centerY (/ yRes 2))
       (r (range 0 359 10))
       )    
    (gimp-undo-push-group-start inImage)
    (while (not (null? r))
	   (drawScale inLayer centerX centerY (car r) 10)
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
    "RGB"                                      ;image type
    SF-IMAGE       "Image"         0
    SF-DRAWABLE    "Drawable"      0
)
(script-fu-menu-register "script-fu-gauge" "<Image>/Filters/Render")

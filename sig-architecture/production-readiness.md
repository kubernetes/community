# Production Readiness Review Process

Production readiness reviews are intended to ensure that features merging into
Kubernetes are observable, scalable and supportable, can be safely operated in
production environments, and can be disabled or rolled back in the event they
cause increased failures in production.

More details may be found in the [PRR KEP].

## Status

As of 1.19, production readiness reviews are required, and are part of the KEP
process. The PRR questionnaire previously found here has been incorporated into
the [KEP template]. The template details the specific questions that must be
answered, depending on the stage of the feature.

As of 1.21, PRRs are now blocking. PRR _approval_ is required for the enhancement
to be part of the release.

Note that some of the questions should be answered in both the KEP's README.md
and the `kep.yaml`, in order to support automated checks on the PRR. The
template points out these as needed.

[PRR KEP]: https://git.k8s.io/enhancements/keps/sig-architecture/1194-prod-readiness
[KEP template]: https://git.k8s.io/enhancements/keps/NNNN-kep-template

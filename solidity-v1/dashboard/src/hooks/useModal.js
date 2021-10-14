import { useCallback, useRef } from "react"
import { useSelector, useDispatch } from "react-redux"
import { modalActions } from "../actions"

export const useModal = () => {
  const modalType = useSelector((state) => state.modal.modalType)
  const modalProps = useSelector((state) => state.modal.modalProps)
  const awaitingPromiseRef = useRef()

  const dispatch = useDispatch()

  const openModal = useCallback(
    (modalType, props) =>
      dispatch(modalActions.showModal({ modalType, modalProps: props })),
    [dispatch]
  )

  const closeModal = useCallback(
    () => dispatch(modalActions.hideModal()),
    [dispatch]
  )

  const onSubmitConfirmationModal = useCallback(
    (values) => {
      if (awaitingPromiseRef.current) {
        awaitingPromiseRef.current.resolve(values)
      }
      closeModal()
    },
    [closeModal]
  )

  const openConfirmationModal = useCallback(
    (modalType, props) => {
      openModal(modalType, { ...props, onConfirm: onSubmitConfirmationModal })

      return new Promise((resolve, reject) => {
        awaitingPromiseRef.current = { resolve, reject }
      })
    },
    [openModal, onSubmitConfirmationModal]
  )

  return {
    modalType,
    modalProps,
    openModal,
    closeModal,
    openConfirmationModal,
  }
}
